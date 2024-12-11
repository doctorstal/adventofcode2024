package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type stone struct {
	val  int64
	next *stone
}

func eleventhDay() {
	fmt.Println("Eleventh day, GO!")
	inputS := "112 1110 163902 0 7656027 83039 9 74"

	head := &stone{}
	curr := head
	for _, s := range strings.Split(inputS, " ") {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		curr.next = &stone{
			val: v,
		}
		curr = curr.next
	}

	for _ = range 25 {
		blinkTransform(head.next)
	}

	fmt.Println(stonesLen(head.next))
}

func blinkTransform(first *stone) {
	for curr := first; curr != nil; curr = curr.next {
		// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
		if curr.val == 0 {
			curr.val = 1
		} else
		// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
		if isEvenDigits(curr.val) {
			p1, p2 := splitStone(curr.val)

			curr.val = p1
			newStone := &stone{
				val:  p2,
				next: curr.next,
			}
			curr.next = newStone
			curr = newStone
		} else
		// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
		{
			curr.val *= 2024
		}
	}
}

func splitStone(num int64) (int64, int64) {
	//l := len(strconv.FormatInt(num, 10)) / 2
	l := int(math.Log10(float64(num))+1) / 2
	p := int64(math.Pow10(l))

	return num / p, num % p
}

func isEvenDigits(num int64) bool {
	//return len(strconv.FormatInt(num, 10))%2 == 0
	return int(math.Log10(float64(num))+1)%2 == 0
}

func printStones(first *stone) {
	for curr := first; curr != nil; curr = curr.next {
		fmt.Print(curr.val, " ")
	}
	fmt.Println()
}

func stonesLen(first *stone) int64 {
	var count int64 = 0
	for curr := first; curr != nil; curr = curr.next {
		count++
	}
	return count
}