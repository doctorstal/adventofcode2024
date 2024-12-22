package main

import (
	"fmt"
	"strconv"
)

func twentyFirstDay() {
	fmt.Println("--- Day 21: Keypad Conundrum ---")
	// codes := readFileAsBytes("input21example.txt")
	codes := readFileAsBytes("input21.txt")

	sumComplexity := 0
	for _, code := range codes {
		sumComplexity += calcComplexity(string(code), 2)
	}

	fmt.Printf("sumComplexity: %v\n", sumComplexity)

	// part2
	sumComplexity25 := 0
	for _, code := range codes {
		sumComplexity25 += calcComplexity(string(code), 24)
	}

	fmt.Printf("sumComplexity25: %v\n", sumComplexity25)
	// 25: 277554934879758 - too high
	// 24: 110880490505014 - too low

}

var keyboardNumPad = [][]byte{
	[]byte("789"),
	[]byte("456"),
	[]byte("123"),
	[]byte(" 0A"),
}

var keyboardDirectional = [][]byte{
	[]byte(" ^A"),
	[]byte("<v>"),
}

func keyCoordinates(keyboard [][]byte) map[byte]Point {
	coords := make(map[byte]Point)
	for i, row := range keyboard {
		for j, v := range row {
			coords[v] = Point{
				y: i,
				x: j,
			}
		}
	}
	return coords
}

func calcComplexity(code string, robotCount int) int {
	numPadCoordinates := keyCoordinates(keyboardNumPad)
	dirPadCoordinates := keyCoordinates(keyboardDirectional)

	input := calcDirectionalInput(map[string]int{code: 1}, numPadCoordinates)

	for _ = range robotCount {
		input = calcDirectionalInput(input, dirPadCoordinates)
	}

	numValue, _ := strconv.ParseInt(code[:3], 10, 32)
	presses := 0
	for s, count := range input {
		presses += count * len(s)
	}

	fmt.Printf("%s: %v, len: %d\n", code, input, presses)

	return presses * int(numValue)
}

func decode(code []byte, keyboard [][]byte) []byte {
	res := make([]byte, 0)
	sy, sx := findOnArea(keyboard, 'A')
	for _, v := range code {
		switch v {
		case 'A':
			res = append(res, keyboard[sy][sx])
		case '>':
			sx++
		case '<':
			sx--
		case 'v':
			sy++
		case '^':
			sy--
		}
		if keyboard[sy][sx] == ' ' {
			fmt.Println("WARNING: Robot will panic!", string(v))
		}
	}
	return res
}

func calcDirectionalInputWithMemo(code map[string]int, keyboard map[byte]Point, memo map[string]map[string]int) map[string]int {

	res := make(map[string]int)
	for key := range code {

		var codedKey map[string]int
		if r, ok := memo[key]; ok {
			codedKey = r
		} else {
			codedKey = calcDirectionalInput(map[string]int{key: 1}, keyboard)
			memo[key] = codedKey
		}
		for k, v := range codedKey {
			res[k] += v
		}
	}
	return res
}

func calcDirectionalInput(code map[string]int, keyboard map[byte]Point) map[string]int {
	appendN := func(r []byte, b byte, n int) []byte {
		for _ = range n {
			r = append(r, b)
		}
		return r
	}

	p := keyboard['A']
	y, x := p.y, p.x

	panikP := keyboard[' ']
	panicY, panicX := panikP.y, panikP.x

	res := make(map[string]int)
	for subcode, count := range code {

		for _, ch := range []byte(subcode) {
			subRes := make([]byte, 0)
			tp := keyboard[ch]
			ty, tx := tp.y, tp.x

			ml := x - tx
			mr := tx - x
			mu := y - ty
			mb := ty - y

			// Avoid panic square
			if ml > 0 && y == panicY && tx == panicX {
				subRes = appendN(subRes, 'v', mb)
				subRes = appendN(subRes, '^', mu)
				mb, mu = 0, 0
			} else if mr > 0 && x == panicX && ty == panicY {
				subRes = appendN(subRes, '>', mr)
				mr = 0
			}
			subRes = appendN(subRes, '<', ml)
			subRes = appendN(subRes, 'v', mb)
			subRes = appendN(subRes, '>', mr)
			subRes = appendN(subRes, '^', mu)

			subRes = append(subRes, 'A')
			y, x = ty, tx
			res[string(subRes)] += count
		}

	}

	return res
}
