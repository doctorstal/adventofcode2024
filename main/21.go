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
		sumComplexity += calcComplexity(code)
	}

	fmt.Printf("sumComplexity: %v\n", sumComplexity)

	
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

func calcComplexity(code []byte) int {
	d1input := calcDirectionalInput(code, keyboardNumPad)
	d2input := calcDirectionalInput(d1input, keyboardDirectional)
	d3input := calcDirectionalInput(d2input, keyboardDirectional)

	numValue, _ := strconv.ParseInt(string(code[:3]), 10, 32)

	fmt.Printf("%s: %s len: %d\n", string(code), string(d3input), len(d3input))

	// validate
	fmt.Println(string(decode(decode(decode(d3input, keyboardDirectional), keyboardDirectional), keyboardNumPad)))
	return len(d3input) * int(numValue)

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

func calcDirectionalInput(code []byte, keyboard [][]byte) []byte {
	appendN := func(r []byte, b byte, n int) []byte {
		for _ = range n {
			r = append(r, b)
		}
		return r
	}
	y, x := findOnArea(keyboard, 'A')

	panicY, panicX := findOnArea(keyboard, ' ')

	res := make([]byte, 0)
	for _, ch := range code {
		ty, tx := findOnArea(keyboard, ch)

		ml := x - tx
		mr := tx - x
		mu := y - ty
		mb := ty - y

		// Avoid panic square
		if ml > 0 && y == panicY && tx == panicX {
			res = appendN(res, 'v', mb)
			res = appendN(res, '^', mu)
			mb, mu = 0, 0
		} else if mr > 0 && x == panicX && ty == panicY {
			res = appendN(res, '>', mr)
			mr = 0
		}
		res = appendN(res, '<', ml)
		res = appendN(res, 'v', mb)
		res = appendN(res, '>', mr)
		res = appendN(res, '^', mu)

		res = append(res, 'A')
		y, x = ty, tx
	}

	return res
}
