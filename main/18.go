package main

import (
	"fmt"
	"strconv"
	"strings"
)

func eighteenthDay() {
	// fmt.Println("--- Day 18: RAM Run ---")
	input := readFileAsBytes("input18.txt")
	const memorySize = 71
	const simulateSteps = 1024
	// input := readFileAsBytes("input18example.txt")
	// const memorySize = 7
	// const simulateSteps = 12

	fallingBytes := make([]*Point, len(input))
	for i, row := range input {
		s := strings.SplitN(string(row), ",", 2)
		p := &Point{}
		p.x, _ = strconv.Atoi(s[0])
		p.y, _ = strconv.Atoi(s[1])
		fallingBytes[i] = p
	}

	//fmt.Printf("parsed: %s\n", fallingBytes)

	memorySpace := make([][]byte, memorySize)
	for i := range memorySize {
		memorySpace[i] = make([]byte, memorySize)
		for j := range memorySize {
			memorySpace[i][j] = '.'
		}
	}
	memorySpace[0][0] = 'S'
	memorySpace[memorySize-1][memorySize-1] = 'E'

	for i := range simulateSteps {
		b := fallingBytes[i]
		memorySpace[b.y][b.x] = '#'
	}

	makeStepsMap := func() [][]int {
		stepsMap := make([][]int, memorySize)
		for i := range memorySize {
			stepsMap[i] = make([]int, memorySize)
			for j := range memorySize {
				stepsMap[i][j] = -1
			}
		}
		return stepsMap
	}

	stepsMap := makeStepsMap()

	markSteps(0, 0, memorySpace, stepsMap, 0)

	printArea(memorySpace)

	fmt.Printf("min steps: %v\n", stepsMap[memorySize-1][memorySize-1])

	// part2
	for i := simulateSteps; i < len(fallingBytes); i++ {
		b := fallingBytes[i]
		memorySpace[b.y][b.x] = '#'

		stepsMap = makeStepsMap()
		markSteps(0, 0, memorySpace, stepsMap, 0)
		if stepsMap[memorySize-1][memorySize-1] == -1 {
			fmt.Printf("last fallen byte: %d,%d\n", b.x, b.y)
			break
		}
	}

}

