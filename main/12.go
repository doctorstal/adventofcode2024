package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func twelfthDay() {
	const size12 = 140

	fmt.Println("Twelfth day go")
	// Read input

	f, err := os.Open("input12.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	plots := make([][]byte, 0, size12)
	for scanner.Scan() {
		plots = append(plots, append(make([]byte, 0, size12), scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	f.Close()

	// *Solution
	// plots [][]byte - map of the plot
	// mapped [][]bool - if plot is already in some region
	// Go through plots, if plot is not mapped - build a region, calculate region cost, mark region as mapped
	// *Build a region
	// Traverse in all directions, until plot is not already mapped and has the same type of plant
	// *Calculate area
	// Add 1 every time new plot in the region is discovered
	// *Calculate perimeter
	// Add 1 for every neighbouring plot of new plot where other type of plant is growing

	mapped := make([][]bool, 0, size12)
	for _ = range size12 {
		mapped = append(mapped, make([]bool, size12))
	}

	var totalCost int64 = 0
	var costWithDiscount int64 = 0

	for i, row := range plots {
		for j, plantType := range row {
			if mapped[i][j] {
				continue
			}
			seenSides := make(map[string]bool)
			area, perimeter, sides := calcAreaAndPerimeter(plots, mapped, seenSides, i, j, plantType)
			totalCost += int64(area) * int64(perimeter)
			costWithDiscount += int64(area) * int64(sides)
		}
	}

	fmt.Printf("TotalCost: %v\n", totalCost)
	fmt.Printf("costWithDiscount: %v\n", costWithDiscount)

}

func isStepOutside(plots [][]byte, i, j int, plantType byte) bool {
	return i < 0 || j < 0 || i >= len(plots) || j >= len(plots) || plots[i][j] != plantType
}

/*
*

	returns area, perimeter, sides
*/
func calcAreaAndPerimeter(plots [][]byte, mapped [][]bool, seenSides map[string]bool, i, j int, plantType byte) (int, int, int) {

	if mapped[i][j] {
		// Ignore already mapped, stop looking
		return 0, 0, 0
	}
	mapped[i][j] = true

	// Plot grows same type of plant and is not mapped before
	area, perimeter, sides := 1, 0, 0

	key := func(a, b, c, d int) string {
		return fmt.Sprintf("%d:%d|%d:%d", a, b, c, d)
	}

	visit := func(di, dj, mi, mj int) {
		var a, p, s int
		if isStepOutside(plots, i+di, j+dj, plantType) {
			// Do not add to area, add to perimeter
			a, p = 0, 1
			if !seenSides[key(i, j, i+di, j+dj)] {
				// mark seen side forward

				// (i, j+1)(di=0, dj=1, mi=1, mj=0) i+m*mi, j+m*mj : i + m * mi + di, j + m * mj + di
				// (i, j-1)(di=0, dj=1, mi=1, mj=0) i+m, j : i+m, j-1
				// (i+1, j)(di=1, dj=0, 0, 1) i, j+m : i+1, j+m
				// (i-1, j)(di=-1, dj=0, 0, 1) i, j+m : i-1, j+m
				for m := 0; m < len(plots); m++ {
					if !isStepOutside(plots, i+m*mi, j+m*mj, plantType) && isStepOutside(plots, i+m*mi+di, j+m*mj+dj, plantType) {
						seenSides[key(i+m*mi, j+m*mj, i+m*mi+di, j+m*mj+dj)] = true
					} else {
						break
					}
				}
				// mark seen side backward

				for m := 0; m > -len(plots); m-- {

					if !isStepOutside(plots, i+m*mi, j+m*mj, plantType) && isStepOutside(plots, i+m*mi+1*di, j+m*mj+1*dj, plantType) {
						seenSides[key(i+m*mi, j+m*mj, i+m*mi+1*di, j+m*mj+1*dj)] = true
					} else {
						break
					}
				}
				s = 1
			}
		} else {
			a, p, s = calcAreaAndPerimeter(plots, mapped, seenSides, i+di, j+dj, plantType)
		}
		area += a
		perimeter += p
		sides += s
	}

	visit(0, 1, 1, 0)
	visit(0, -1, 1, 0)
	visit(1, 0, 0, 1)
	visit(-1, 0, 0, 1)

	return area, perimeter, sides

}
