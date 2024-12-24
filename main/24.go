package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
)

type Gate struct {
	w1       string
	operator string
	w2       string
	wout     string
}

func (g *Gate) Out(wires map[string]byte) (out byte, ok bool) {
	v1, v1Ok := wires[g.w1]
	v2, v2Ok := wires[g.w2]

	if v1Ok && v2Ok {
		switch g.operator {
		case "OR":
			return v1 | v2, true
		case "AND":
			return v1 & v2, true
		case "XOR":
			return v1 ^ v2, true
		default:
			panic("Unknown operator!")
		}

	}
	return 0, false
}

func twentyFourthDay() {
	fmt.Println("--- Day 24: Crossed Wires ---")
	wires, gates := readWiresAndGates("input/input24.txt")
	// wires, gates := readWiresAndGates("input24example.txt")

	fmt.Printf("wires: %v\n", wires)
	fmt.Printf("gates: %s\n", gates)

	res := processGates(wires, gates)
	fmt.Printf("res: %b, %d\n", res, res)
}

func processGates(wires map[string]byte, gates []*Gate) int64 {
	zWires := findZWires(gates)
	fmt.Printf("zWires: %v\n", zWires)

	res := int64(0)
	hasRes := false
	for ; !hasRes; res, hasRes = findOutputOnZWires(wires, zWires) {
		for _, g := range gates {
			if gOut, ok := g.Out(wires); ok {
				wires[g.wout] = gOut
			}
		}

	}

	return res
}

func findOutputOnZWires(wires map[string]byte, zWires []string) (out int64, ok bool) {
	for _, zw := range zWires {
		out <<= 1
		if b, ok := wires[zw]; ok {
			fmt.Print(int64(b))
			out += int64(b)
		} else {
			return 0, false
		}
	}
	return out, true

}

func findZWires(gates []*Gate) []string {
	zWires := make([]string, 0)
	for _, g := range gates {
		if g.wout[0] == 'z' {
			zWires = append(zWires, g.wout)
		}
	}
	slices.Sort(zWires)
	slices.Reverse(zWires)
	return zWires
}

func readWiresAndGates(filename string) (wires map[string]byte, gates []*Gate) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wires = make(map[string]byte)

	wirePattern := regexp.MustCompile(`(\w+): (\d)`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		submatches := wirePattern.FindSubmatch(scanner.Bytes())
		wires[string(submatches[1])] = submatches[2][0] - '0'

	}

	gates = make([]*Gate, 0)
	gatePattern := regexp.MustCompile(`(\w+) (\w+) (\w+) -> (\w+)`)

	for scanner.Scan() {
		submatches := gatePattern.FindStringSubmatch(scanner.Text())
		gates = append(gates, &Gate{
			w1:       submatches[1],
			operator: submatches[2],
			w2:       submatches[3],
			wout:     submatches[4],
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wires, gates

}
