package main

import "fmt"

func twentiethDay() {
	fmt.Println("--- Day 20: Race Condition ---")

	// area := readFileAsBytes("input20example.txt")
	area := readFileAsBytes("input20.txt")

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
	skipCount := findShortcutsCount(steps, minSkip, 2, area)
	skipCount20 := findShortcutsCount(steps, minSkip, 20, area)

	fmt.Printf("findShortcutsCount(steps, %d): %v\n", minSkip, skipCount)
	fmt.Printf("findShortcutsCount20(steps, %d): %v\n", minSkip, skipCount20)
}

func findShortcutsCount(steps [][]int, minSkip int, skipLength int, area [][]byte) int {
	/*
		2:
		#######
		###2###
		#######
		#2#S#2#
		#######
		###2###
		#######

		3:
		###3###
		##323##
		#31113#
		321S123
		#31113#
		##323##
		###3###
	*/

	h, w := len(steps), len(steps[0])
	skipVal := func(sy, sx, ey, ex int) int {
		if ey >= 0 && ey < h && ex >= 0 && ex < w && steps[ey][ex] != -1 {
			return steps[ey][ex] - steps[sy][sx]
		}
		return -1
	}

	skipInc := func(sy, sx, ey, ex int, sl int) int {
		if sv := skipVal(sy, sx, ey, ex); sv >= minSkip+sl {
			return 1
		}
		return 0
	}

	skipsCount := 0

	for i := 1; i < len(steps); i++ {
		for j := 1; j < len(steps[0]); j++ {
			if steps[i][j] != -1 {
				for k := -skipLength; k <= skipLength; k++ {
					for l := -skipLength + abs(k); l <= skipLength-abs(k); l++ {
						skipsCount += skipInc(i, j, i+k, j+l, abs(k)+abs(l))
					}
				}

			}
		}
	}
	return skipsCount
}
