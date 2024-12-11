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

	rules := make([][]int, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rule := make([]int, 0, 2)
		for _, v := range strings.SplitN(scanner.Text(), "|", 2) {
			parsed, _ := strconv.Atoi(v)
			rule = append(rule, parsed)
		}
		rules = append(rules, rule)
	}
	updates := make([][]int, 0)
	for scanner.Scan() {
		update := make([]int, 0)
		for v := range strings.Split(scanner.Text(), ",") {
			update = append(update, v)
		}
		updates = append(updates, update)
	}

	fmt.Println(rules, updates)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
