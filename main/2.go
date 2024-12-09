package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func secondDay() {
	f, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	safeCount := 0
	safeP2Count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isSafe(line) {
			safeCount++
		} else if isSafeP2(line) {
			safeP2Count++
			//println(line)
		}
		// else {
		// 	println(line)
		// }
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(safeCount)
	println(safeP2Count)
	println(safeP2Count + safeCount)

	println(isSafeP2("26 28 25 23 20 20"))
}

func isSafe(line string) bool {
	reportS := strings.Split(line, " ")

	report := make([]int, 0, len(reportS))
	for _, s := range reportS {
		v, _ := strconv.Atoi(s)
		report = append(report, v)
	}

	prev := report[0]
	firstDiff := report[0] - report[1]

	for i := 1; i < len(report); i++ {
		step := prev - report[i]
		if firstDiff < 0 {
			step = -step
		}
		if step < 1 || step > 3 {
			return false
		}
		prev = report[i]
	}

	return true
}

func isSafeP2(line string) bool {
	reportS := strings.Split(line, " ")

	report := make([]int, 0, len(reportS))
	for _, s := range reportS {
		v, _ := strconv.Atoi(s)
		report = append(report, v)
	}

	for i := range report {
		if isSafeP2Skip(report, i) {
			return true
		}
	}

	return false
}

func isSafeP2Skip(report []int, skip int) bool {
	fi := 0
	si := 1
	if skip == 0 {
		fi = 1
		si = 2
	} else if skip == 1 {
		fi = 0
		si = 2
	}

	prev := report[fi]
	firstDiff := report[fi] - report[si]
	fmt.Println(firstDiff, fi, si)

	for i := si; i < len(report); i++ {
		if i == skip {
			continue
		}
		step := prev - report[i]
		if firstDiff < 0 {
			step = -step
		}
		if step < 1 || step > 3 {

			//printReport(report, skip, i, step)

			return false
		}
		prev = report[i]
	}

	return true
}
