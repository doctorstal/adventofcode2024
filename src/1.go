package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	firstList := make([]int, 0, 100)
	secondList := make([]int, 0, 100)

	for scanner.Scan() {
		pair := strings.SplitN(scanner.Text(), "   ", 2)

		if val, err := strconv.Atoi(pair[0]); err == nil {
			firstList = append(firstList, val)
		}

		if val, err := strconv.Atoi(pair[1]); err == nil {
			secondList = append(secondList, val)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	diff := 0

	sort.Ints(firstList)
	sort.Ints(secondList)

	for i := range firstList {
		a := firstList[i]
		b := secondList[i]
		if a > b {
			diff += a - b
		} else {
			diff += b - a
		}
	}

	// part2

	pos1 := 0
	pos2 := 0
	simmilarity := 0

	for pos1 < len(firstList) && pos2 < len(secondList) {
		num := firstList[pos1]

		count1 := 0
		count2 := 0

		for pos1 < len(firstList) && num == firstList[pos1] {
			count1++
			pos1++
		}

		for pos2 < len(secondList) && num > secondList[pos2] {
			pos2++
		}
		for pos2 < len(secondList) && num == secondList[pos2] {
			count2++
			pos2++
		}

		if count2 > 0 {
			fmt.Println(num, count1, count2)
		}

		simmilarity += num * count1 * count2
	}

	fmt.Println(diff)
	fmt.Println(simmilarity)
}
