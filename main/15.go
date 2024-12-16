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

	wharehouse, moves := readWharehouseAndMoves("input15.txt")

	fmt.Printf("wharehouse: \n %s\n", bytes.Join(wharehouse, []byte{'\n'}))
	fmt.Printf("moves: \n %s\n", moves)

	wideWH := widenWharehouse(wharehouse)

	whRobot := findWhRobot(wharehouse)

	wharehouse[whRobot.y][whRobot.x] = '.'

	fmt.Printf("whRobot: %v\n", whRobot)

	for _, m := range moves {
		move(whRobot, wharehouse, m)
	}
	printWharehouse(wharehouse, whRobot)
	fmt.Println("GPS coordinates:", calcGPSCoordinates(wharehouse))

	// p2
	fmt.Printf("wideWH: \n%s\n", bytes.Join(wideWH, []byte{'\n'}))

	whRobot = findWhRobot(wideWH)
	wideWH[whRobot.y][whRobot.x] = '.'
	for _, m := range moves {
		move(whRobot, wideWH, m)
		//fmt.Printf("Move: %s\n", string(m))
		//printWharehouse(wideWH, whRobot)
	}

	printWharehouse(wideWH, whRobot)
	fmt.Println("GPS coordinates p2:", calcGPSCoordinates(wideWH))

}

func widenWharehouse(wharehouse [][]byte) [][]byte {
	wideWH := make([][]byte, 0, len(wharehouse))
	for _, row := range wharehouse {
		wideRow := make([]byte, 0, len(row))

		for _, tile := range row {
			switch tile {
			case '.':
				wideRow = append(wideRow, '.', '.')
			case 'O':
				wideRow = append(wideRow, '[', ']')
			case '#':
				wideRow = append(wideRow, '#', '#')
			case '@':
				wideRow = append(wideRow, '@', '.')
			default:
				panic(fmt.Sprintf("Illegal character in input %s", string(tile)))
			}
		}
		wideWH = append(wideWH, wideRow)
	}
	return wideWH
}

func calcGPSCoordinates(wharehouse [][]byte) (res int) {
	for i, row := range wharehouse {
		for j, tile := range row {
			if tile == 'O' {
				res += i*100 + j
			}
			// p2
			if tile == '[' {
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

	tile := wharehouse[robot.y+dy][robot.x+dx]
	switch tile {
	case '.':
		robot.x += dx
		robot.y += dy
	case '#':
		// stay in place
		return
	case 'O':
		// Following logic does not cover wide boxes, but we do not care, as they could not be in the same room
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
	case '[', ']':
		// Following logic does not cover narrow boxes, but we do not care, as they could not be in the same room

		boxesToMove := make([][]int, 1)
		boxesToMove[0] = make([]int, 2)
		boxesToMove[0][0], boxesToMove[0][1] = robot.y+dy, robot.x+dx
		if tile == ']' {
			// covers left and misaligned vertical pushes
			boxesToMove[0][1]--
		}

		for i := 0; i < len(boxesToMove); i++ {
			box := boxesToMove[i]

			tx, ty := box[1]+dx, box[0]+dy
			targetTile1 := wharehouse[ty][tx]
			targetTile2 := wharehouse[ty][tx+1]
			if targetTile1 == '#' || targetTile2 == '#' {
				// we hit the wall, nothing can move
				return
			} else if targetTile1 == '[' {
				// it's a vertical push of single aligned box
				boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx))
			} else if targetTile1 == ']' {
				if targetTile2 == '[' {
					if dx == 1 {
						boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx+1))
					} else if dx == -1 {
						boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx-1))
					} else {
						// vertical push of two boxes
						boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx-1))
						boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx+1))
					}
				} else if targetTile2 == '.' {
					if dx == 0 {
						// vertical push of one misaligned box
						boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx-1))
					}
				}
			} else if targetTile1 == '.' && targetTile2 == '[' {
				if dx == 0 {
					// vertical push of one misaligned box
					boxesToMove = append(boxesToMove, append(make([]int, 0, 2), ty, tx+1))
				}
			}
		}

		// move boxes
		// 1. clear old boxes positions
		for _, box := range boxesToMove {
			wharehouse[box[0]][box[1]] = '.'
			wharehouse[box[0]][box[1]+1] = '.'
		}
		// 2. put boexs in new positions
		for _, box := range boxesToMove {
			tx, ty := box[1]+dx, box[0]+dy
			wharehouse[ty][tx] = '['
			wharehouse[ty][tx+1] = ']'
		}
		robot.x += dx
		robot.y += dy

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
