package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fourthDay() {
	fmt.Println("Fourth day go!")

	size := 140

	f, err := os.Open("input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([][]byte, 0, size)

	for scanner.Scan() {
		b := append(make([]byte, 0, size), scanner.Bytes()...)
		input = append(input, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum := 0
			sum += findWordCount("XMAS", input, i, j)
			sum += findWordCount("SAMX", input, i, j)
			if sum > 0 {
				fmt.Print(string(input[i][j]))
			} else {
				fmt.Print(".")
			}

			res += sum
		}
		fmt.Println()
	}

	fmt.Println("XMAS count: ", res)

	res2 := 0

	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			if isX_Mas(input, i, j) {
				res2++
				fmt.Print("A")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("X-MAS count:", res2)
}

func isX_Mas(input [][]byte, i, j int) bool {
	isA := input[i][j] == 'A'
	isM_Sd1 := input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S'
	isS_Md1 := input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M'

	isM_Sd2 := input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S'
	isS_Md2 := input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M'

	return isA && (isM_Sd1 || isS_Md1) && (isM_Sd2 || isS_Md2)
}

func findWordCount(word string, arr [][]byte, i int, j int) int {
	h, v, d1, d2 := 1, 1, 1, 1
	for p := 0; p < len(word); p++ {
		c := word[p]
		if (j+p) >= len(arr[i]) || arr[i][j+p] != c {
			h *= 0
		}
		if (i+p) >= len(arr) || arr[i+p][j] != c {
			v *= 0
		}
		if (j+p) >= len(arr[i]) || (i+p) >= len(arr) || arr[i+p][j+p] != c {
			d1 *= 0
		}
		if (j-p) < 0 || (i+p) >= len(arr) || arr[i+p][j-p] != c {
			d2 *= 0
		}
	}

	return h + v + d1 + d2
}
