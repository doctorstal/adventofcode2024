package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const size12 = 140

func twelfthDay() {
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

	for i, row := range plots {
		for j, plantType := range row {
			if mapped[i][j] {
				continue
			}
			area, perimeter := calcAreaAndPerimeter(plots, mapped, i, j, plantType)
			totalCost += int64(area) * int64(perimeter)
		}
	}

	fmt.Printf("TotalCost: %v\n", totalCost)

}

/*
*

	returns area, perimeter
*/
func calcAreaAndPerimeter(plots [][]byte, mapped [][]bool, i, j int, plantType byte) (int, int) {
	if i < 0 || j < 0 || i >= len(plots) || j >= len(plots) || plots[i][j] != plantType {
		// Do not add to area, add to perimeter, stop looking
		return 0, 1
	}
	if mapped[i][j] {
		// Ignore already mapped, stop looking
		return 0, 0
	}
	mapped[i][j] = true

	// Plot grows same type of plant and is not mapped before
	area, perimeter := 1, 0

	var a, p int
	// look right
	a, p = calcAreaAndPerimeter(plots, mapped, i+1, j, plantType)
	area += a
	perimeter += p

	// look left
	a, p = calcAreaAndPerimeter(plots, mapped, i-1, j, plantType)
	area += a
	perimeter += p

	// look up
	a, p = calcAreaAndPerimeter(plots, mapped, i, j-1, plantType)
	area += a
	perimeter += p

	// look down
	a, p = calcAreaAndPerimeter(plots, mapped, i, j+1, plantType)
	area += a
	perimeter += p

	return area, perimeter

}
