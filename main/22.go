package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func twentySecondDay() {
	fmt.Println("--- Day 22: Monkey Market ---")
	// prices := readMonkeyPrices("input22example.txt")
	prices := readMonkeyPrices("input22.txt")

	var pricesSum int64 = 0
	for _, n := range prices {
		pricesSum += calc2000thPrice(n)
	}

	fmt.Printf("pricesSum: %v\n", pricesSum)

	fmt.Printf("findMostBananas(prices): %v\n", findMostBananas(prices))
	// fmt.Printf("findMostBananas(prices): %v\n", findMostBananas([]int64{123}))
	//fmt.Printf("findMostBananas(prices): %v\n", findMostBananas([]int64{123}))

}

func calc2000thPrice(n int64) int64 {
	for _ = range 2000 {
		n = nextMonkeySecret(n)
	}
	return n
}

func nextMonkeySecret(n int64) int64 {
	n = mixIntoSecret(n, n*64)
	n = pruneSecret(n)
	n = mixIntoSecret(n, n/32)
	n = pruneSecret(n)

	n = mixIntoSecret(n, n*2048)
	n = pruneSecret(n)
	return n
}

func mixIntoSecret(n int64, r int64) int64 {
	return n ^ r
}
func pruneSecret(n int64) int64 {
	return n % 16777216
}

func findMostBananas(secrets []int64) int64 {
	prices := make(map[string]int64)

	for _, secret := range secrets {
		for k, p := range buildSeqPrices(secret) {
			prices[k] += p
		}
	}

	maxK, maxP := "", int64(0)

	for k, p := range prices {
		if maxP < p {
			maxP = p
			maxK = k
		}
	}

	fmt.Printf("maxK: %v\n", maxK)

	return maxP
}
func buildSeqPrices(n int64) map[string]int64 {
	sw := NewSlidingWindow(4)

	res := make(map[string]int64)
	price := int64(n % 10)
	diff := int64(0)

	// for i := range 2000 {
	for i := range 1999 {
		n = nextMonkeySecret(n)

		prevPrice := price
		price = n % 10
		diff = price - prevPrice
		sw.Push(diff)

		// fmt.Println(n, "\t:", price, diff)

		if i > 4 {
			key := sw.State()
			// set if absent
			if _, ok := res[key]; !ok {
				res[key] = price
			}
		}

	}
	return res
}

func readMonkeyPrices(filename string) []int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	res := make([]int64, 0)
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, n)
	}
	return res
}

type SlidingWindow struct {
	size int
	idx  int
	data []int64
}

func NewSlidingWindow(size int) *SlidingWindow {
	sw := &SlidingWindow{
		size: size,
		idx:  0,
		data: make([]int64, size),
	}
	return sw
}

func (sw *SlidingWindow) Push(el int64) {
	sw.data[sw.idx] = el
	sw.idx = (sw.idx + 1) % sw.size
}

func (sw *SlidingWindow) State() string {
	s := make([]string, sw.size)
	for i := 0; i < sw.size; i++ {
		idx := (sw.idx + i) % sw.size
		s[i] = strconv.FormatInt(int64(sw.data[idx]), 10)
	}
	return strings.Join(s, ",")
}
