package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Section [2]int

func (s1 Section) Contains(s2 Section) bool {
	return s1[0] <= s2[0] && s1[1] >= s2[1]
}

func (s1 Section) Overlaps(s2 Section) bool {
	if s1[0] < s2[0] {
		return s1[1] >= s2[0]
	}
	return s2[1] >= s1[0]
}

type Pair struct {
	First, Second Section
}

func fullyContainedPairs(pairs []Pair) (sum int) {
	for _, p := range pairs {
		if p.First.Contains(p.Second) || p.Second.Contains(p.First) {
			sum++
		}
	}
	return
}

func overlappingPairs(pairs []Pair) (sum int) {
	for _, p := range pairs {
		if p.First.Overlaps(p.Second) {
			sum++
		}
	}
	return
}

func parseInput(f io.Reader) []Pair {
	var pairs []Pair
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var pair Pair
		r := strings.NewReader(scanner.Text())
		fmt.Fscanf(r, "%d-%d,%d-%d\n",
			&pair.First[0], &pair.First[1],
			&pair.Second[0], &pair.Second[1])
		pairs = append(pairs, pair)
	}

	return pairs
}
