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

func (guard *Guard) getGuardByte() byte {
	if guard.dy == -1 {
		return '^'
	}
	if guard.dy == 1 {
		return 'v'
	}
	if guard.dx == -1 {
		return '<'
	}
	if guard.dx == 1 {
		return '>'
	}
	return 'X'
}

func (guard *Guard) getPathKey() string {
	return fmt.Sprintf("%d,%d,%s", guard.x, guard.y, string(guard.getGuardByte()))
}

func (guard *Guard) moveInArea(area [][]byte, aw, ah int) bool {
	guard.x += guard.dx
	guard.y += guard.dy
	if guard.x < 0 || guard.x >= aw || guard.y < 0 || guard.y >= ah {
		return false
	}

	if area[guard.y][guard.x] == '#' {
		// step back
		guard.x -= guard.dx
		guard.y -= guard.dy
		// rotate 90
		guard.dy, guard.dx = guard.dx, -guard.dy
	}
	return true
}

func sixthDay() {
	fmt.Println("Day 6: Guard Gallivant")

	area := readMappedArea("input6.txt")
	// area := readMappedArea("input6example.txt")
	guard := findGuard(area)

	fmt.Printf("guard: %v\n", *guard)

	ah := len(area)
	aw := len(area[0])

	obstacleOptions := 0
	path := make(map[string]bool)
	for {
		if canLoop(area, guard, path) {
			obstacleOptions++
		}
		path[guard.getPathKey()] = true
		area[guard.y][guard.x] = guard.getGuardByte()

		if !guard.moveInArea(area, aw, ah) {
			break
		}
	}

	positions := 0

	for _, row := range area {
		for _, tile := range row {
			if tile != '.' && tile != '#' {
				positions++
			}
		}
		fmt.Println(string(row))
	}

	fmt.Printf("positions: %v\n", positions)
	fmt.Printf("obstacleOptions: %v\n", obstacleOptions)

}

func canLoop(area [][]byte, guard *Guard, path map[string]bool) bool {
	if area[guard.y][guard.x] != '.' {
		return false
	}

	area[guard.y][guard.x] = '#'

	res := false
	testGuard := &Guard{
		x:  guard.x - guard.dx,
		y:  guard.y - guard.dy,
		dx: guard.dx,
		dy: guard.dy,
	}
	ah := len(area)
	aw := len(area[0])

	testPath := make(map[string]bool)

	for {
		if !testGuard.moveInArea(area, aw, ah) {
			res = false
			break
		}

		key := testGuard.getPathKey()

		if path[key] || testPath[key] {
			res = true
			break
		}
		testPath[key] = true
	}

	area[guard.y][guard.x] = '.'
	return res
}

func findGuard(area [][]byte) *Guard {
	for y, row := range area {
		for x, tile := range row {
			if tile == '^' {
				return &Guard{
					x:  x,
					y:  y,
					dx: 0,
					dy: -1,
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
