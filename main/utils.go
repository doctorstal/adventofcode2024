package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type (
	Point struct {
		y int
		x int
	}
	DirectedPoint struct {
		Point
		dx int
		dy int
	}
)

func readFileAsBytes(fileName string) (area [][]byte) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	area = make([][]byte, 0)
	for scanner.Scan() {
		area = append(area, append(make([]byte, 0), scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return area
}

func printArea(area [][]byte) {
	fmt.Printf("area:\n%s\n", bytes.Join(area, []byte{'\n'}))
}

func findOnArea(area [][]byte, target byte) (y, x int) {
	for i, row := range area {
		for j, tile := range row {
			if tile == target {
				return i, j
			}
		}
	}
	return -1, -1
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
