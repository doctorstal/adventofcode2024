package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ninteenthDay() {
	fmt.Println("--- Day 19: Linen Layout ---")

	// suply, designs := readLinenSuplyAndDesigns("input19.txt")
	suply, designs := readLinenSuplyAndDesigns("input19.txt")

	//fmt.Println(suply, designs)

	possibleCount := 0
	memoPossible := make(map[string]bool)
	for _, design := range designs {
		if designPossible(design, suply, memoPossible) {
			possibleCount++
		}
	}
	fmt.Printf("possibleCount: %v\n", possibleCount)

	possibleWaysCount := 0
	memoWays := make(map[string]int)
	for _, design := range designs {
		possibleWaysCount += countPossibleWays(design, suply, memoWays)
	}
	fmt.Printf("possibleWaysCount: %v\n", possibleWaysCount)

}

func designPossible(design string, suply []string, memo map[string]bool) bool {
	if res, ok := memo[design]; ok {
		return res
	}
	if design == "" {
		return true
	}

	for _, s := range suply {
		nd, hasPrefix := strings.CutPrefix(design, s)
		if hasPrefix && designPossible(nd, suply, memo) {
			memo[design] = true
			return true
		}
	}
	memo[design] = false
	return false
}

func countPossibleWays(design string, suply []string, memo map[string]int) int {
	if res, ok := memo[design]; ok {
		return res
	}
	if design == "" {
		return 1
	}

	count := 0

	for _, s := range suply {
		nd, hasPrefix := strings.CutPrefix(design, s)
		if hasPrefix {
			count += countPossibleWays(nd, suply, memo)
		}
	}
	memo[design] = count
	return count
}

func readLinenSuplyAndDesigns(fileName string) (suply, designs []string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if scanner.Scan() {
		suply = strings.Split(scanner.Text(), ", ")
	}
	scanner.Scan() // skip empty line
	designs = make([]string, 0)
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return suply, designs
}
