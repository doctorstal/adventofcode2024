package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type WHRobot struct {
	x int
	y int
}

func fifteenthDay() {
	fmt.Println("--- Day 15: Warehouse Woes ---")

	wharehouse, moves := readWharehouseAndMoves("input15example.txt")

	fmt.Printf("wharehouse: \n %s\n", bytes.Join(wharehouse, []byte{'\n'}))
	fmt.Printf("moves: \n %s\n", moves)

	whRobot := findWhRobot(wharehouse)

	wharehouse[whRobot.y][whRobot.y] = '.'

	fmt.Printf("whRobot: %v\n", whRobot)

	for _, m := range moves {
		move(whRobot, wharehouse, m)
	}
	printWharehouse(wharehouse, whRobot)
	fmt.Println("GPS coordinates:", calcGPSCoordinates(wharehouse))
}

func calcGPSCoordinates(wharehouse [][]byte) (res int) {
	for i, row := range wharehouse {
		for j, tile := range row {
			if tile == 'O' {
				res += i*100 + j
			}
		}
	}
	return res
}

func move(robot *WHRobot, wharehouse [][]byte, move byte) {
	dx, dy := 0, 0
	switch move {
	case '^':
		dy = -1
	case 'v':
		dy = 1
	case '>':
		dx = 1
	case '<':
		dx = -1
	default:
		panic("Invalid move!")
	}

	switch wharehouse[robot.y+dy][robot.x+dx] {
	case '.':
		robot.x += dx
		robot.y += dy
	case '#':
		// stay in place
	case 'O':
		sy, sx := robot.y+dy, robot.x+dx
		for wharehouse[sy][sx] == 'O' {
			sx += dx
			sy += dy
		}
		if wharehouse[sy][sx] == '.' {
			wharehouse[sy][sx] = 'O'
			wharehouse[robot.y+dy][robot.x+dx] = '.'
			robot.x += dx
			robot.y += dy
		}
	}

}

func printWharehouse(wharehouse [][]byte, robot *WHRobot) {
	for i, row := range wharehouse {
		for j, tile := range row {
			if i == robot.y && j == robot.x {
				fmt.Print("@")
			} else {
				fmt.Print(string(tile))
			}
		}
		fmt.Println()
	}
}

func findWhRobot(wharehouse [][]byte) *WHRobot {
	for y, row := range wharehouse {
		for x, tile := range row {
			if tile == '@' {
				return &WHRobot{
					x: x,
					y: y,
				}
			}
		}
	}

	return &WHRobot{
		x: -1,
		y: -1,
	}
}

func readWharehouseAndMoves(fileName string) (wharehouse [][]byte, moves []byte) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	wharehouse = make([][]byte, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		wharehouse = append(wharehouse, append(make([]byte, 0), scanner.Bytes()...))
	}
	moves = make([]byte, 0)
	for scanner.Scan() {
		moves = append(moves, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wharehouse, moves
}
