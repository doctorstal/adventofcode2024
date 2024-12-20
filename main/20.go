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
	// skipCount := findShortcutsCount(steps, minSkip, 2, area)
	skipCount20 := findShortcutsCount(steps, minSkip, 20, area)

	printArea(area)

	// fmt.Printf("findShortcutsCount(steps, %d): %v\n", minSkip, skipCount)
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

		i <- -3 to 3
			j <- -3+abs(i) to 3-abs(i)

		i <- -x to x
			j <- -x+abs(i) to x-abs(i)
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
			// area[sy][sx] = '1'
			// area[sy][sx] = '1'
			// area[ey][ex] = '2'
			// area[ey][ex] = '2'
			return 1
		}
		return 0
	}

	skipsCount := 0

	for i := 1; i < len(steps); i++ {
		for j := 1; j < len(steps[0]); j++ {
			if steps[i][j] != -1 {
				// k <- -x to x
				// 		l <- -x+abs(k) to x-abs(k)
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
