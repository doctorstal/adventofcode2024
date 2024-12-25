package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func twentyFifthDay() {
	fmt.Println("--- Day 25: Code Chronicle ---")

	// locks, keys := readLocksAndKeys("input25exapmle.txt")
	locks, keys := readLocksAndKeys("input/input25.txt")

	// fmt.Printf("locks: %v\n", locks)
	// fmt.Printf("keys: %v\n", keys)

	fitCount := 0
	for _, k := range keys {
		for _, l := range locks {
			if k.fitsLock(l) {
				fitCount++
				// fmt.Println("Fit lock/key pair", l, k)
			}
		}
	}
	fmt.Printf("fitCount: %v\n", fitCount)

}

type Lock [5]int
type Key [5]int

func (k *Key) fitsLock(l Lock) bool {
	for i := range k {
		if k[i] > l[i] {
			return false
		}
	}
	return true
}

func readLocksAndKeys(filename string) (locks []Lock, keys []Key) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	locks = make([]Lock, 0)
	keys = make([]Key, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "....." {

			key := Key{}
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					keys = append(keys, key)
					break
				}

				for i, c := range line {
					if c == '#' {
						key[i]++
					} else if key[i] > 0 {
						fmt.Println("Warning! Hole in key!")
					}
				}
			}
		} else if scanner.Text() == "#####" {

			lock := Lock{}
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					locks = append(locks, lock)
					break
				}

				for i, c := range line {
					if c == '.' {
						lock[i]++
					} else if lock[i] > 0 {
						fmt.Println("Warning! Hole in lock!")
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return locks, keys
}
