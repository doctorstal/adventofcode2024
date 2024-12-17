package main

import (
	"fmt"
)

type (
	MazePoint struct {
		DirectedPoint
		score int
	}
)

func sixteenthDay() {
	fmt.Println("--- Day 16: Reindeer Maze ---")
	maze := readArea("input16.txt")
	printArea(maze)

	fmt.Println("Min score:", findMazePathScore(maze))
}

func findMazePathScore(maze [][]byte) int {
	sy, sx := findOnArea(maze, 'S')
	scores := make([][]int, len(maze))
	for i, row := range maze {
		scores[i] = make([]int, len(row))
	}

	marcPathScore(maze, sy, sx-1, 0, 1, scores, 0)
	// for _, row := range scores {
	// 	fmt.Println(row)
	// }

	ey, ex := findOnArea(maze, 'E')

	return scores[ey][ex]

}

func marcPathScore(maze [][]byte, y, x, dy, dx int, scores [][]int, currScore int) {
	cy, cx := y+dy, x+dx
	tile := maze[cy][cx]
	// fmt.Println("cx", cx, "cy", cy, "tile", string(tile), "score", scores[cy][cx])
	switch tile {
	case '#':
		return
	case '.', 'S':
		if scores[cy][cx] == 0 || scores[cy][cx] > currScore {
			scores[cy][cx] = currScore
			marcPathScore(maze, cy, cx, dy, dx, scores, currScore+1)
			// turn right
			marcPathScore(maze, cy, cx, dx, -dy, scores, currScore+1001)
			// turn left
			marcPathScore(maze, cy, cx, -dx, dy, scores, currScore+1001)
			// do not turn back
		}
	case 'E':
		if scores[cy][cx] == 0 || scores[cy][cx] > currScore {
			scores[cy][cx] = currScore
		}
	}
}
