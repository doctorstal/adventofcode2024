package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fifthDay() {
	fmt.Println("Day Five Go, Go, Go!")
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

	fmt.Println(rules, updates)

	res := 0

	for _, update := range updates {
		if validateUpdateOrder(update, rules) {
			res += update[len(update)/2]
		}
	}

	fmt.Printf("res: %v\n", res)
}

func validateUpdateOrder(update []int, rules map[int][]int) bool {
	fmt.Println(update)
	for i, v := range update {
		for j := i - 1; j >= 0; j-- {
			if bRule, ok := rules[v]; ok {
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
