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
	maze := readFileAsBytes("input16example.txt") //57349
	printArea(maze)

	score, tiles := findMazePathScoreAndBestTiles(maze)

	fmt.Println("Min score:", score)
	fmt.Printf("tiles: %v\n", tiles)
}

func findMazePathScoreAndBestTiles(maze [][]byte) (score, tiles int) {
	sy, sx := findOnArea(maze, 'S')
	scores := make([][]map[string]int, len(maze))
	for i, row := range maze {
		scores[i] = make([]map[string]int, len(row))
		for j := range row {
			scores[i][j] = make(map[string]int)
		}
	}

	path := make(map[string]bool)

	marcPathScore(maze, sy, sx-1, 0, 1, scores, 0, path)

	ey, ex := findOnArea(maze, 'E')

	for i, row := range maze {
		for j := range row {
			if path[fmt.Sprintf("%d:%d", i, j)] {
				maze[i][j] = 'O'
			}
		}
	}
	printArea(maze)

	score = 0
	for _, v := range scores[ey][ex] {
		if score == 0 || score > v {
			score = v
		}
	}

	fmt.Printf("pathsFound: %v\n", pathsFound)

	return score, len(path)

}

var pathsFound = 0

func marcPathScore(maze [][]byte, y, x, dy, dx int, scores [][]map[string]int, currScore int, path map[string]bool) (found bool) {
	cy, cx := y+dy, x+dx
	key := fmt.Sprintf("%d:%d", cy, cx)
	dKey := fmt.Sprintf("%d:%d", dy, dx)
	tile := maze[cy][cx]
	// fmt.Println("cx", cx, "cy", cy, "tile", string(tile), "score", scores[cy][cx])
	switch tile {
	case '#':
		return false
	case '.', 'S':
		if scores[cy][cx][dKey] == 0 || scores[cy][cx][dKey] >= currScore {
			scores[cy][cx][dKey] = currScore

			if marcPathScore(maze, cy, cx, dy, dx, scores, currScore+1, path) {
				path[key] = true
			}
			// turn right
			if marcPathScore(maze, cy, cx, dx, -dy, scores, currScore+1001, path) {
				path[key] = true
			}
			// turn left
			if marcPathScore(maze, cy, cx, -dx, dy, scores, currScore+1001, path) {
				path[key] = true
			}
			return path[key]
			// do not turn back
		}
	case 'E':
		score := 0
		for _, v := range scores[cy][cx] {
			if score == 0 || score > v {
				score = v
			}
		}
		if score == 0 || score > currScore {
			scores[cy][cx][dKey] = currScore
			for k := range path {
				delete(path, k)
			}
			path[key] = true
			pathsFound = 1
			println("clear, found")
			return true
		}
		if score == currScore {
			println("found")
			pathsFound++
			return true
		}
	}
	return false
}
