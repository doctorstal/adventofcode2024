package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const size10 = 54

func tenthDay() {

	fmt.Println("Tenth Day Go!")
	f, err := os.Open("input10.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	input := make([][]byte, 0, size10)

	for scanner.Scan() {
		input = append(input,
			append(make([]byte, 0, size10), scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	totalScore := 0
	totalRaiting := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i][j] == '0' {
				score := countTrailheadScore(input, i, j, newSeen())
				totalScore += score
				totalRaiting += countTrailheadRaiting(input, i, j)
			}
		}
	}

	fmt.Println("Sum of scores: ", totalScore)
	fmt.Println("Sum of raitings: ", totalRaiting)
}

func newSeen() [][]bool {
	seen := make([][]bool, 0, size10)
	for i := 0; i < size10; i++ {
		seen = append(seen, make([]bool, size10))
	}
	return seen
}

func countTrailheadScore(input [][]byte, i, j int, seen [][]bool) int {
	if input[i][j] == '9' {
		if seen[i][j] {
			return 0
		} else {
			seen[i][j] = true
			return 1
		}
	}

	score := 0
	nextAlt := input[i][j] + 1
	if j+1 < len(input[i]) && input[i][j+1] == nextAlt {
		score += countTrailheadScore(input, i, j+1, seen)
	}
	if i+1 < len(input) && input[i+1][j] == nextAlt {
		score += countTrailheadScore(input, i+1, j, seen)
	}
	if j >= 1 && input[i][j-1] == nextAlt {
		score += countTrailheadScore(input, i, j-1, seen)
	}
	if i >= 1 && input[i-1][j] == nextAlt {
		score += countTrailheadScore(input, i-1, j, seen)
	}
	return score
}

func countTrailheadRaiting(input [][]byte, i, j int) int {
	if input[i][j] == '9' {
		return 1
	}

	score := 0
	nextAlt := input[i][j] + 1
	if j+1 < len(input[i]) && input[i][j+1] == nextAlt {
		score += countTrailheadRaiting(input, i, j+1)
	}
	if i+1 < len(input) && input[i+1][j] == nextAlt {
		score += countTrailheadRaiting(input, i+1, j)
	}
	if j >= 1 && input[i][j-1] == nextAlt {
		score += countTrailheadRaiting(input, i, j-1)
	}
	if i >= 1 && input[i-1][j] == nextAlt {
		score += countTrailheadRaiting(input, i-1, j)
	}
	return score
}
