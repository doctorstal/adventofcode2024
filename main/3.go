package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func thirdDay() {
	fmt.Println("Third day go!")
	f, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := make([]byte, 0)
	for scanner.Scan() {
		input = append(input, scanner.Bytes()...)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	inputS := string(input)

	fmt.Println(calcMuls(inputS))

	// part2

	fmt.Println(calcMuls(filterDonts(inputS)))

}

func filterDonts(input string) string {
	pattern := regexp.MustCompile(`don't\(\).*?do\(\)`)
	return pattern.ReplaceAllString(input, "")
}

func calcMuls(inputS string) int64 {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	mults := pattern.FindAllStringSubmatch(inputS, -1)

	var res int64 = 0

	for _, v := range mults {
		a, _ := strconv.Atoi(v[1])
		b, _ := strconv.Atoi(v[2])
		res += int64(a * b)
	}
	return res
}
