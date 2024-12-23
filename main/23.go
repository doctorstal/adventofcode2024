package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

func twentyThirdDay() {
	fmt.Println("--- Day 23: LAN Party ---")
	// networkConnectivity, networkLinks := readNetworkMap("input23example.txt")
	networkConnectivity, networkLinks := readNetworkMap("input23.txt")

	fmt.Printf("sets: %v\n", countTrippleSets(networkConnectivity, networkLinks))

	fmt.Printf("Party: %s\n", findLargestNetParty(networkConnectivity, networkLinks))
}

type NetParty map[string]bool

func partyPass(np map[string]bool) string {
	return strings.Join(slices.Sorted(maps.Keys(np)), ",")
}

func findLargestNetParty(connectivity map[string]map[string]bool, links map[string][]string) string {
	found := make(map[string]bool)

	passesToSize := make(map[string]int)

	isInParty := func(c string, party NetParty) bool {
		for cp := range party {
			if !connectivity[c][cp] {
				return false
			}
		}
		return true

	}

	for c1, neighbours := range links {
		if found[c1] {
			continue
		}
		for len(neighbours) > 0 {
			toAddNext := make([]string, 0)
			party := NetParty{c1: true}
			for _, c2 := range neighbours {
				if isInParty(c2, party) {
					found[c2] = true
					party[c2] = true
				} else {
					toAddNext = append(toAddNext, c2)
				}
			}
			passesToSize[partyPass(party)] = len(party)
			neighbours = toAddNext
		}

	}

	maxSize, pass := 0, ""
	for p, s := range passesToSize {
		if maxSize < s {
			maxSize = s
			pass = p
		}
	}

	return pass
}

func countTrippleSets(networkConnectivity map[string]map[string]bool, networkLinks map[string][]string) (setsCount int) {
	for c1, links := range networkLinks {
		for i := 0; i < len(links); i++ {
			c2 := links[i]
			for j := i + 1; j < len(links); j++ {
				c3 := links[j]
				if networkConnectivity[c2][c3] && (c1[0] == 't' || c2[0] == 't' || c3[0] == 't') {
					setsCount++
				}
			}
		}
	}
	return setsCount / 3
}

func readNetworkMap(filename string) (connectivity map[string]map[string]bool, links map[string][]string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	connectivity = make(map[string]map[string]bool)
	links = make(map[string][]string)
	addLink := func(c1, c2 string) {
		if _, ok := connectivity[c1]; !ok {
			connectivity[c1] = make(map[string]bool, 1)
			links[c1] = make([]string, 0)
		}
		connectivity[c1][c2] = true
		links[c1] = append(links[c1], c2)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		link := strings.SplitN(scanner.Text(), "-", 2)
		addLink(link[0], link[1])
		addLink(link[1], link[0])

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return connectivity, links

}
