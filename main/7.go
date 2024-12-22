package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Equation struct {
	testValue int
	numbers   []int
}

func (eq *Equation) isValidEquation() bool {
	return isValidEquation(eq.testValue, 0, eq.numbers)
}

func isValidEquation(target, sum int, numbers []int) bool {
	if len(numbers) == 0 {
		return target == sum
	}
	head := numbers[0]
	tail := numbers[1:]
	return isValidEquation(target, sum+head, tail) || isValidEquation(target, sum*head, tail)

}

func seventhDay() {
	fmt.Println("--- Day 7: Bridge Repair ---")
	// equations := readEquations("input7.txt")
	equations := readEquations("input7example.txt")

	fmt.Printf("equations: %s\n", equations)

	validSum := 0
	for _, eq := range equations {
		if eq.isValidEquation() {
			validSum += eq.testValue
		}
	}

	fmt.Printf("validSum: %v\n", validSum)

}

func readEquations(filename string) []*Equation {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	testValuePattern := regexp.MustCompile(`^(\d+):`)
	numbersPattern := regexp.MustCompile(` (\d+)`)
	equations := make([]*Equation, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		eq := &Equation{}
		s := scanner.Text()

		eq.testValue, _ = strconv.Atoi(testValuePattern.FindStringSubmatch(s)[1])

		submatches := numbersPattern.FindAllStringSubmatch(s, -1)

		eq.numbers = make([]int, len(submatches))
		for i := 0; i < len(submatches); i++ {
			eq.numbers[i], _ = strconv.Atoi(submatches[i][1])
		}
		equations = append(equations, eq)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return equations
}