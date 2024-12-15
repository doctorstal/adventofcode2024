package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Guard struct {
	x  int
	y  int
	dx int
	dy int
}

func sixthDay() {
	fmt.Println("Day 6: Guard Gallivant")

	area := readMappedArea("input6.txt")
	guard := findGuard(area)

	fmt.Printf("guard: %v\n", *guard)

	ah := len(area)
	aw := len(area[0])
	for {
		area[guard.y][guard.x] = 'X'

		guard.x += guard.dx
		guard.y += guard.dy
		if guard.x < 0 || guard.x >= aw || guard.y < 0 || guard.y >= ah {
			break
		}

		if area[guard.y][guard.x] == '#' {
			// step back
			guard.x -= guard.dx
			guard.y -= guard.dy
			// rotate 90
			guard.dy, guard.dx = guard.dx, -guard.dy
		}
	}

	positions := 0

	for _, row := range area {
		for _, tile := range row {
			if tile == 'X' {
				positions++
			}
		}
		fmt.Println(string(row))
	}

	fmt.Printf("positions: %v\n", positions)

}

func findGuard(area [][]byte) *Guard {
	for y, row := range area {
		for x, tile := range row {
			dx, dy := 0, 0
			switch tile {
			case '^':
				dy = -1
			case 'v':
				dy = 1
			case '>':
				dx = 1
			case '<':
				dx = -1

			default:
				continue
			}
			if tile == '^' {
				return &Guard{
					x:  x,
					y:  y,
					dx: dx,
					dy: dy,
				}
			}
		}
	}

	return &Guard{
		x: -1,
		y: -1,
	}
}

func readMappedArea(fileName string) [][]byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	area := make([][]byte, 0)
	for scanner.Scan() {
		area = append(area, append(make([]byte, 0), scanner.Bytes()...))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return area

}
