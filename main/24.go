package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
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

type FullAdder struct {
	//in wires
	cin string
	x   string
	y   string
	// out wires
	z    string
	cout string
	// gates
	sGate    *Gate // 	x XOR y s
	zGate    *Gate //     cin XOR s z
	c0Gate   *Gate //     cin AND s c0
	c1Gate   *Gate // 	x AND y c1
	coutGate *Gate //      c0 OR c1 cout!
}

func twentyFourthDay() {
	fmt.Println("--- Day 24: Crossed Wires ---")
	// 51745744348272
	// Extected 27165461756787+24670290535861=
	// 51835752292648
	wires, gates := readWiresAndGates("input/input24.txt")
	// wires, gates := readWiresAndGates("input24example.txt")

	res := processGates(wires, gates)
	fmt.Printf("res: %b, %d\n", res, res)

	// part2
	gatesByOut := make(map[string]*Gate)
	for _, g := range gates {
		gatesByOut[g.wout] = g
	}

	// I've checked z0 and z1 manually, as z0 is half-adder

	adders := make([]*FullAdder, 0, 45)

	for i := 2; i < 45; i++ {
		key := fmt.Sprintf("z%02d", i)
		fmt.Printf("key: %v\n", key)

		zGate := gatesByOut[key]

		if zGate.operator != "XOR" {
			log.Println("Z gate has wrong operator", zGate)
			continue
		}
		g1, g2 := gatesByOut[zGate.w1], gatesByOut[zGate.w2]
		var sGate *Gate
		var cinGate *Gate
		if g1.operator == "OR" && g2.operator == "XOR" {
			cinGate, sGate = g1, g2
		} else if g1.operator == "XOR" && g2.operator == "OR" {
			cinGate, sGate = g2, g1
		} else {
			log.Println("s or cin Gates has wrong operators", zGate, g1, g2)
			continue
		}

		s := sGate.wout
		cin := cinGate.wout
		var c0Gate *Gate

		if g, ok := findGate(gates, s, cin, "AND"); ok {
			c0Gate = g
		} else {
			log.Println("c1Gate not found!")
			continue
		}

		c0 := c0Gate.wout

		x, y := sGate.w1, sGate.w2
		var c1Gate *Gate

		if g, ok := findGate(gates, x, y, "AND"); ok {
			c1Gate = g
		} else {
			log.Println("c1Gate not found!")
			continue
		}

		c1 := c1Gate.wout

		var coutGate *Gate

		if g, ok := findGate(gates, c1, c0, "OR"); ok {
			coutGate = g
		} else {
			log.Println("coutGate not found!")
			continue
		}

		cout := coutGate.wout

		adder := &FullAdder{
			cin: cin,
			x:   sGate.w1,
			y:   sGate.w2,

			z:    zGate.wout,
			cout: cout,

			sGate:    sGate,    // 	x XOR y s
			zGate:    zGate,    //     cin XOR s z
			c0Gate:   c0Gate,   //     cin AND s c0
			c1Gate:   c1Gate,   // 	x AND y c1
			coutGate: coutGate, //      c0 OR c1 cout!
		}
		adders = append(adders, adder)
	}

	for i := 1; i < len(adders); i++ {
		if adders[i].cin != adders[i-1].cout {
			log.Println("cin does not match", adders[i-1].z, adders[i-1].cout, adders[i].cin)
		}
	}

	for _, a := range adders {
		delete(gatesByOut, a.sGate.wout)
		delete(gatesByOut, a.zGate.wout)
		delete(gatesByOut, a.c0Gate.wout)
		delete(gatesByOut, a.c1Gate.wout)
		delete(gatesByOut, a.coutGate.wout)
	}

	for _, g := range gatesByOut {
		fmt.Printf("%s %s %s -> %s\n", g.w1, g.operator, g.w2, g.wout)
	}

	outsToSwap := []string{"hmt", "z18", "bfq", "z27", "hkh", "z31", "fjp", "bng"}

	slices.Sort(outsToSwap)
	fmt.Println(strings.Join(outsToSwap, ","))

	// sort.Slice(gates, func(i, j int) bool {
	// 	return gates[i].w1 > gates[j].w1
	// })

	// for i, g := range gates {
	// 	fmt.Printf("%d: %s %s %s %s\n", i, g.w1, g.operator, g.w2, g.wout)
	// }
}

func findGate(gates []*Gate, w1, w2, operator string) (*Gate, bool) {
	for _, g := range gates {
		if ((g.w1 == w1 && g.w2 == w2) || (g.w1 == w2 && g.w2 == w1)) && g.operator == operator {
			return g, true
		}
	}
	return nil, false
}

func processGates(wires map[string]byte, gates []*Gate) int64 {
	zWires := findWires(gates, 'z')

	xWires := make([]string, 0, 44)
	yWires := make([]string, 0, 44)

	for i := 0; i < 45; i++ {
		xWires = append(xWires, fmt.Sprintf("x%02d", i))
		yWires = append(yWires, fmt.Sprintf("y%02d", i))
	}

	x, _ := findOutputOnWires(wires, xWires)
	y, _ := findOutputOnWires(wires, yWires)

	fmt.Printf("Extected %d+%d=%d \n", x, y, x+y)

	fmt.Printf("xWires: %v\n", xWires)
	fmt.Printf("yWires: %v\n", yWires)

	res := int64(0)
	hasRes := false
	for ; !hasRes; res, hasRes = findOutputOnWires(wires, zWires) {
		for _, g := range gates {
			if _, ok := wires[g.wout]; ok {
				continue
			}
			if gOut, ok := g.Out(wires); ok {
				wires[g.wout] = gOut
			}
		}

	}

	return res
}

func findOutputOnWires(wires map[string]byte, zWires []string) (out int64, ok bool) {
	for _, zw := range zWires {
		out <<= 1
		if b, ok := wires[zw]; ok {
			out += int64(b)
		} else {
			return 0, false
		}
	}
	return out, true

}

func findWires(gates []*Gate, firstByte byte) []string {
	zWires := make([]string, 0)
	for _, g := range gates {
		if g.wout[0] == firstByte {
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
		w1w2 := []string{
			submatches[1],
			submatches[3],
		}
		slices.Sort(w1w2)
		gates = append(gates, &Gate{
			w1:       w1w2[0],
			operator: submatches[2],
			w2:       w1w2[1],
			wout:     submatches[4],
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wires, gates

}
