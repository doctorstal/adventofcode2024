package main

import "fmt"

func eighthDay() {
	fmt.Println("--- Day 8: Resonant Collinearity ---")
	area := readFileAsBytes("input8example.txt")
	printArea(area)

	antenas := make(map[byte][]*Point)

	for i, row := range area {
		for j, tile := range row {
			if tile != '.' {
				p := &Point{
					y: i,
					x: j,
				}
				if points, ok := antenas[tile]; ok {
					antenas[tile] = append(points, p)
				} else {
					antenas[tile] = []*Point{p}
				}
			}
		}
	}

	ah, aw := len(area), len(area[0])

	antinodes := make(map[Point]bool)

	for _, points := range antenas {

		for i, p1 := range points {
			for j := i + 1; j < len(points); j++ {
				p2 := points[j]
				addAntinode(p1, p2, antinodes, ah, aw)
				addAntinode(p2, p1, antinodes, ah, aw)
			}
		}

	}

	fmt.Println("========")

	for i, row := range area {
		for j, tile := range row {
			if antinodes[Point{y: i, x: j}] {
				fmt.Print(string('#'))
			} else {
				fmt.Print(string(tile))
			}
		}
		fmt.Println()
	}

	fmt.Printf("antinodes: %v\n", antinodes)
	fmt.Printf("len(antinodes): %v\n", len(antinodes))
	//fmt.Printf("antenas: %s\n", antenas)

}

func addAntinode(p1, p2 *Point, antinodes map[Point]bool, ah, aw int) {
	an := Point{
		y: 2*p1.y - p2.y,
		x: 2*p1.x - p2.x,
	}
	if an.x >= 0 && an.x < aw && an.y >= 0 && an.y < ah {
		antinodes[an] = true
	}
}
