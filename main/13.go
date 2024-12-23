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
	ax int64
	ay int64
	bx int64
	by int64
	px int64
	py int64
}

func thirteenthDay() {
	fmt.Println("Thirtinth day go!")

	f, err := os.Open("input/input13.txt")
	// f, err := os.Open("input13example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	machines := readMachines(f)

	fmt.Println("Tokens needed: ", calcTokensForAllPrizes(machines, 0))

	// Part 2

	fmt.Println("Tokens needed for real: ", calcTokensForAllPrizes(machines, 10000000000000))

	// findDifferenceP1AndP2(machines)

}

func calcTokensForAllPrizes(machines []*ClawMachine, d int64) int64 {
	res := int64(0)
	for _, m := range machines {
		m.px += d
		m.py += d

		res += calcTokensForPrizeMath(m)
	}
	return res
}

func calcTokensForPrizeMath(m *ClawMachine) int64 {
	// A = (p_x*b_y - prize_y*b_x) / (a_x*b_y - a_y*b_x)
	// B = (a_x*p_y - a_y*p_x) / (a_x*b_y - a_y*b_x)
	d := m.ax*m.by - m.ay*m.bx
	a := (m.px*m.by - m.py*m.bx) / d
	b := (m.py*m.ax - m.px*m.ay) / d

	if (a*m.ax+b*m.bx == m.px) && (a*m.ay+b*m.by == m.py) {
		return 3*a + b
	} else {
		return 0
	}
}

func calcTokensForPrizeIter(m *ClawMachine) int64 {

	maxB := min(m.px/m.bx, m.py/m.by)
	fmt.Println(maxB)
	x, y := m.bx*maxB, m.by*maxB
	for b := maxB; b > 0; b-- {
		a := (m.px - x) / m.ax
		if m.px == x+a*m.ax && m.py == a*m.ay+y {
			fmt.Println(a, b, maxB)
			return a*3 + b*1
		}
		x -= m.bx
		y -= m.by
	}
	return 0

}

/*
	func calcTokensForAllPrizes(machines []*ClawMachine) int {
		res := 0
		for _, m := range machines {
			t, _ := calcTokensForPrize(m, 0, 0, 0, 0, make(map[string]int))
			res += t
		}
		return res
	}

	func calcTokensForPrize(m *ClawMachine, x, y, pushNumA, pushNumB int, memo map[string]int) (tokens int, ok bool) {
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
*/
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

			machine.ax, _ = strconv.ParseInt(baMatches[1], 10, 64)
			machine.ay, _ = strconv.ParseInt(baMatches[2], 10, 64)

			machine.bx, _ = strconv.ParseInt(bbMatches[1], 10, 64)
			machine.by, _ = strconv.ParseInt(bbMatches[2], 10, 64)
			machine.px, _ = strconv.ParseInt(pMatches[1], 10, 64)
			machine.py, _ = strconv.ParseInt(pMatches[2], 10, 64)

			machines = append(machines, machine)

		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return machines
}
