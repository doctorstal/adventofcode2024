package main

import "fmt"

func twentiethDay() {
	fmt.Println("--- Day 20: Race Condition ---")

	// area := readFileAsBytes("input20example.txt")
	area := readFileAsBytes("input20.txt")
	printArea(area)

	sy, sx := findOnArea(area, 'S')

	steps := make([][]int, len(area))
	for i, row := range area {
		steps[i] = make([]int, len(row))
		for j := range row {
			steps[i][j] = -1
		}
	}

	markSteps(sy, sx, area, steps, 0)

	minSkip := 100
	skipCount := findShortcutsCount(steps, minSkip, area)

	printArea(area)

	fmt.Printf("findShortcutsCount(steps, %d): %v\n", minSkip, skipCount)
}

func findShortcutsCount(steps [][]int, minSkip int, area [][]byte) int {
	h, w := len(steps), len(steps[0])
	skipVal := func(sy, sx, ey, ex int) int {
		if ey >= 0 && ey < h && ex >= 0 && ex < w && steps[ey][ex] != -1 {
			return steps[ey][ex] - steps[sy][sx]
		}
		return -1
	}

	skipInc := func(sy, sx, ey, ex int) int {
		if sv := skipVal(sy, sx, ey, ex); sv >= minSkip+2 {
			area[sy][sx] = '1'
			area[sy][sx] = '1'
			area[ey][ex] = '2'
			area[ey][ex] = '2'
			//fmt.Println(sy, sx, ey, ex, sv)
			return 1
		}
		return 0
	}

	skipsCount := 0
	for i := 1; i < len(steps); i++ {
		for j := 1; j < len(steps[0]); j++ {
			if steps[i][j] != -1 {
				skipsCount += skipInc(i, j, i+2, j)
				skipsCount += skipInc(i, j, i-2, j)
				skipsCount += skipInc(i, j, i, j+2)
				skipsCount += skipInc(i, j, i, j-2)
			}
		}
	}
	return skipsCount
}


