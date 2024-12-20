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
	maze := readFileAsBytes("input16.txt")
	printArea(maze)

	score, tiles := findMazePathScoreAndBestTiles(maze)

	fmt.Println("Min score:", score)
	fmt.Printf("tiles: %v\n", tiles)
}

func printFoundPath(path map[string]bool) {
	maze := readFileAsBytes("input16.txt")
	for i, row := range maze {
		for j := range row {
			if path[fmt.Sprintf("%d:%d", i, j)] {
				maze[i][j] = ' '
			}
		}
	}
	printArea(maze)
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

	markPathScore(maze, sy, sx-1, 0, 1, scores, 0, path)

	ey, ex := findOnArea(maze, 'E')

	score = 0
	for _, v := range scores[ey][ex] {
		if score == 0 || score > v {
			score = v
		}
	}

	fmt.Printf("pathsFound: %v\n", pathsFound)

	//printFoundPath(bestTiles)

	return score, len(bestTiles)

}

var pathsFound = 0
var bestTiles map[string]bool

func markPathScore(maze [][]byte, y, x, dy, dx int, scores [][]map[string]int, currScore int, path map[string]bool) {
	cy, cx := y+dy, x+dx
	key := fmt.Sprintf("%d:%d", cy, cx)
	dKey := fmt.Sprintf("%d:%d", dy, dx)
	tile := maze[cy][cx]

	if path[key] {
		return
	}

	switch tile {
	case '#':
		return
	case '.', 'S':
		if scores[cy][cx][dKey] == 0 || scores[cy][cx][dKey] >= currScore {
			scores[cy][cx][dKey] = currScore

			path[key] = true

			markPathScore(maze, cy, cx, dy, dx, scores, currScore+1, path)
			// turn right
			markPathScore(maze, cy, cx, dx, -dy, scores, currScore+1001, path)
			// turn left
			markPathScore(maze, cy, cx, -dx, dy, scores, currScore+1001, path)

			delete(path, key)

			//fmt.Printf("len(path): %v : %s\n", len(path), key)

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

			bestTiles = make(map[string]bool)

			for k := range path {
				bestTiles[k] = true
			}
			bestTiles[key] = true

			pathsFound = 1
			println("clear, found", len(path))
			return
		}
		if score == currScore {
			println("found more", len(path))
			for k := range path {
				bestTiles[k] = true
			}
			pathsFound++
			return
		}
	}
}
