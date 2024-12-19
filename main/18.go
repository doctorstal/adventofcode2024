package main

import (
	"fmt"
	"strconv"
	"strings"
)

func eighteenthDay() {
	fmt.Println("--- Day 18: RAM Run ---")
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

	stepsMap := make([][]int, memorySize)
	for i := range memorySize {
		stepsMap[i] = make([]int, memorySize)
		for j := range memorySize {
			stepsMap[i][j] = -1
		}
	}

	markSteps(0, 0, memorySpace, stepsMap, 0)

	printArea(memorySpace)

	fmt.Printf("min steps: %v\n", stepsMap[memorySize-1][memorySize-1])

}

func markSteps(y, x int, space [][]byte, stepsMap [][]int, stepsCount int) {
	if y < 0 || y >= len(space) || x < 0 || x >= len(space[0]) || space[y][x] == '#' {
		return
	}
	if stepsMap[y][x] == -1 || stepsMap[y][x] > stepsCount {
		stepsMap[y][x] = stepsCount
		markSteps(y+1, x, space, stepsMap, stepsCount+1)
		markSteps(y, x+1, space, stepsMap, stepsCount+1)
		markSteps(y-1, x, space, stepsMap, stepsCount+1)
		markSteps(y, x-1, space, stepsMap, stepsCount+1)
	}
}
