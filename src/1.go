package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
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
	
	res := 0
	
	sort.Ints(firstList)
	sort.Ints(secondList)
	
	for i := range(firstList) {
		a := firstList[i]
		b := secondList[i]
		if (a > b) {
			res += a - b
		} else {
			res += b - a
		}
	}

	fmt.Println(res)
}
