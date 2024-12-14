package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func fourteenthDay() {
	spaceW, spaceH := 101, 103
	//spaceW, spaceH := 11, 7
	simulateSeconds := 100

	fmt.Println("Fourteenth day Go!")
	f, err := os.Open("input14.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots := make([]*Robot, 0)

	inputPattern := regexp.MustCompile(`p=([\-0-9]+),([\-0-9]+) v=([\-0-9]+),([\-0-9]+)`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := inputPattern.FindStringSubmatch(scanner.Text())
		robot := &Robot{}

		robot.x, _ = strconv.Atoi(matches[1])
		robot.y, _ = strconv.Atoi(matches[2])
		robot.vx, _ = strconv.Atoi(matches[3])
		robot.vy, _ = strconv.Atoi(matches[4])

		robots = append(robots, robot)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	seconds := 0

	for _ = range simulateSeconds {
		simulateRobotMovement(robots, spaceW, spaceH)
		seconds++
	}

	fmt.Println("SafetyFactor:", calcSafetyFactor(robots, spaceW, spaceH))

	// part2

	reader := bufio.NewReader(os.Stdin)
	for i := simulateSeconds; i > 0; i++ {
		simulateRobotMovement(robots, spaceW, spaceH)
		seconds++
		if findTreeMaybe(robots, spaceW, spaceH) {

			fmt.Println("seconds ", seconds)
			printRobots(robots, spaceW, spaceH)

			if s, _ := reader.ReadString('\n'); s == "end" {
				break
			}
		}
	}
}

func calcSafetyFactor(robots []*Robot, spaceW, spaceH int) int {
	var qtl, qtr, qbl, qbr int
	for _, r := range robots {
		if r.x < (spaceW / 2) {
			if r.y < spaceH/2 {
				qtl++
			} else if r.y > spaceH/2 {
				qbl++
			}
		} else if r.x > spaceW/2 {
			if r.y < spaceH/2 {
				qtr++
			} else if r.y > spaceH/2 {
				qbr++
			}
		}
	}
	fmt.Println(qtl, qtr, qbl, qbr)
	return qtl * qtr * qbl * qbr
}

func simulateRobotMovement(robots []*Robot, spaceW, spaceH int) {
	for _, r := range robots {
		r.x = (r.x + spaceW + r.vx) % spaceW
		r.y = (r.y + spaceH + r.vy) % spaceH

	}
}

func findTreeMaybe(robots []*Robot, spaceW, spaceH int) bool {
	robotCount := make(map[int]int)
	for _, r := range robots {
		robotCount[r.x*spaceW+r.y]++
	}
	robotsInRow := 0
	for i := range spaceH {
		for j := range spaceW {
			if _, ok := robotCount[i+j*spaceW]; ok {
				robotsInRow++
				if robotsInRow > 10 {
					return true
				}
			} else {
				robotsInRow = 0
			}
		}

	}
	return false
}

func printRobots(robots []*Robot, spaceW, spaceH int) {
	for _ = range spaceW {
		fmt.Print("=")
	}
	fmt.Println()

	robotCount := make(map[int]int)
	for _, r := range robots {
		robotCount[r.x*spaceW+r.y]++
	}
	for i := range spaceH {
		for j := range spaceW {
			if c, ok := robotCount[i+j*spaceW]; ok {
				fmt.Print(c)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
