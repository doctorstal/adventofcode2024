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

func fifthDay() {
	fmt.Println("Day Five Print Queue Go!")
	f, err := os.Open("input5.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[int][]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rule := make([]int, 0, 2)
		for _, v := range strings.SplitN(scanner.Text(), "|", 2) {
			parsed, _ := strconv.Atoi(v)
			rule = append(rule, parsed)
		}
		if _, ok := rules[rule[0]]; !ok {
			rules[rule[0]] = make([]int, 0)
		}
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	updates := make([][]int, 0)
	for scanner.Scan() {
		update := make([]int, 0)
		for _, v := range strings.Split(scanner.Text(), ",") {
			parsed, _ := strconv.Atoi(v)
			update = append(update, parsed)
		}
		updates = append(updates, update)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := 0
	corr := 0

	for _, update := range updates {
		if validateUpdateOrder(update, rules) {
			res += update[len(update)/2]
		} else {
			corrected := correctUpdateOrder(update, rules)
			corr += corrected[len(corrected)/2]

		}
	}

	fmt.Printf("res: %v\n", res)
	fmt.Printf("corr: %v\n", corr)
}

func correctUpdateOrder(update []int, rules map[int][]int) []int {
	corrected := append(make([]int, 0, len(update)), update...)
	sort.Slice(corrected, func(i, j int) bool {
		// [i] before [j]?
		if bRule, ok := rules[corrected[j]]; ok {
			for _, b := range bRule {
				if corrected[i] == b {
					return false
				}
			}
		}
		return true
	})

	return corrected
}

func validateUpdateOrder(update []int, rules map[int][]int) bool {

	for i, v := range update {
		if bRule, ok := rules[v]; ok {
			for j := i - 1; j >= 0; j-- {
				for _, b := range bRule {
					if update[j] == b {
						return false
					}
				}
			}
		}
	}

	return true

}
