package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ClawMachine struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

func thirteenthDay() {
	fmt.Println("Thirtinth day go!")

	f, err := os.Open("input13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	machines := readMachines(f)

	fmt.Println("Tokens needed: ", calcTokensForAllPrizes(machines))

}

func calcTokensForAllPrizes(machines []*ClawMachine) int {
	res := 0
	for _, m := range machines {
		memo := make(map[string]int)
		t, _ := calcTokensForPrize(m, 0, 0, 0, 0, memo)
		res += t
	}
	return res
}

func calcTokensForPrize(m *ClawMachine, x, y, pushNumA int, pushNumB int, memo map[string]int) (tokens int, ok bool) {
	// aCost := 3
	// bCost := 1
	//maxPushes := 100

	memoKey := fmt.Sprintf("%d:%d", x, y)

	if c, ok := memo[memoKey]; ok {
		// probably not an ideomatic way for Golang
		if c == -1 {
			return 0, false
		} else {
			return c, true
		}

	}
	if pushNumA >= 100 || pushNumB >= 100 {
		//memo[memoKey] = -1
		return 0, false
	}
	if m.px == x && m.py == y {
		memo[memoKey] = 0
		return 0, true
	}

	ta, okA := calcTokensForPrize(m, x+m.ax, y+m.ay, pushNumA+1, pushNumB, memo)
	tb, okB := calcTokensForPrize(m, x+m.bx, y+m.by, pushNumA, pushNumB+1, memo)

	ta += 3
	tb += 1

	if okA && okB {
		if ta > tb {
			memo[memoKey] = tb
			return tb, true
		} else {
			memo[memoKey] = ta
			return ta, true
		}
	} else if okA {
		memo[memoKey] = ta
		return ta, true
	} else if okB {
		memo[memoKey] = tb
		return tb, true
	} else {
		memo[memoKey] = -1
		return 0, false
	}

}

func readMachines(f *os.File) []*ClawMachine {
	scanner := bufio.NewScanner(f)

	buttonAPattern := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBPattern := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizePattern := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	clawMachineBuffer := [4]string{}
	buffPos := 0

	machines := make([]*ClawMachine, 0)

	for scanner.Scan() {
		clawMachineBuffer[buffPos] = scanner.Text()
		buffPos++
		if buffPos >= 4 {
			buffPos = 0

			baMatches := buttonAPattern.FindStringSubmatch(clawMachineBuffer[0])
			bbMatches := buttonBPattern.FindStringSubmatch(clawMachineBuffer[1])
			pMatches := prizePattern.FindStringSubmatch(clawMachineBuffer[2])

			machine := &ClawMachine{}

			machine.ax, _ = strconv.Atoi(baMatches[1])
			machine.ay, _ = strconv.Atoi(baMatches[2])

			machine.bx, _ = strconv.Atoi(bbMatches[1])
			machine.by, _ = strconv.Atoi(bbMatches[2])

			machine.px, _ = strconv.Atoi(pMatches[1])
			machine.py, _ = strconv.Atoi(pMatches[2])

			machines = append(machines, machine)

		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return machines
}
