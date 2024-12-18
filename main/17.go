package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ChronospatialComputer struct {
	iptr int
	out  []int
	regA int
	regB int
	regC int
}

func (c *ChronospatialComputer) runProgram(program []int) (out string) {
	c.iptr = 0
	c.out = make([]int, 0)
	for c.iptr < len(program) {
		opcode := program[c.iptr]
		operand := program[c.iptr+1]
		c.iptr += c.performOperation(opcode, operand)
	}

	output := ""
	for _, v := range c.out {
		temp := strconv.Itoa(v)
		output = output + temp + ","
	}

	return output
}
func (c *ChronospatialComputer) getComboOperandValue(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	default:
		panic("Invalid combo operand!")
	}
}

func (c *ChronospatialComputer) performOperation(opcode, operand int) (jump int) {
	switch opcode {
	// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
	case 0:
		c.regA = c.regA / (1 << c.getComboOperandValue(operand))
	// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
	case 1:
		c.regB = c.regB ^ operand
		// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
	case 2:
		c.regB = c.getComboOperandValue(operand) % 8
		// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
	case 3:
		if c.regA != 0 {
			c.iptr = operand
			return 0
		}
		// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
	case 4:
		c.regB = c.regB ^ c.regC
		// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)
	case 5:
		c.out = append(c.out, c.getComboOperandValue(operand)%8)
		// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)
	case 6:
		c.regB = c.regA / (1 << c.getComboOperandValue(operand))
		// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)
	case 7:
		c.regC = c.regA / (1 << c.getComboOperandValue(operand))
	}
	return 2
}

func seventeenthDay() {
	fmt.Println("--- Day 17: Chronospatial Computer ---")

	c := &ChronospatialComputer{}

	input := readFileAsBytes("input17.txt")
	// input := readFileAsBytes("input17example.txt")

	reqPattern := regexp.MustCompile(`: (\d+)`)
	c.regA, _ = strconv.Atoi(reqPattern.FindStringSubmatch(string(input[0]))[1])
	c.regB, _ = strconv.Atoi(reqPattern.FindStringSubmatch(string(input[1]))[1])
	c.regC, _ = strconv.Atoi(reqPattern.FindStringSubmatch(string(input[2]))[1])

	programPattern := regexp.MustCompile(`: ([,0-9]+)`)
	programString := programPattern.FindStringSubmatch(string(input[4]))[1]
	program := make([]int, 0)
	for _, s := range strings.Split(programString, ",") {
		n, _ := strconv.Atoi(s)
		program = append(program, n)
	}

	fmt.Println(c, program)

	fmt.Printf("c.runProgram(program): %v\n", c.runProgram(program))
	fmt.Printf("c: %v\n", c)

	matchLast := func(a1, a2 []int) bool {
		for j := 0; j < len(c.out); j++ {
			if a1[j] != a2[len(a2)-len(a1)+j] {
				return false
			}
		}
		return true
	}

	bCount := 0

	for i := 0; i < math.MaxInt; i++ {
		c.regA, c.regB, c.regC = i, 0, 0
		c.runProgram(program)

		if matchLast(c.out, program) {
			fmt.Printf("c.runProgram(program): %d: %v\n", i, c.out)
			if len(c.out) == len(program) {
				fmt.Printf("FOUND! %d\n", i)
				return
			}

			// Go down the branch
			i = i<<3 - 1
			bCount = 0

		} else {
			bCount++
		}
		if bCount > 1024 {
			// loop back to prev tree branch
			i = i >> 3
		}
		//fmt.Printf("i: %b\n", i)
	}

}
