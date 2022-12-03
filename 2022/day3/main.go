package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := parseInput(string(f))

	fmt.Println(sumPriorityChars(input))
	fmt.Println(sumBadges(input))
}

func parseInput(input string) (parsed [][2]string) {
	split := strings.Split(input, "\n")
	for _, s := range split {
		if len(s) > 1 {
			l := len(s) / 2
			parsed = append(parsed, [2]string{s[:l], s[l:]})
		}
	}
	return
}

func findCommonChar(given [2]string) string {
	first := given[0]
	second := given[1]
	for _, c1 := range strings.Split(first, "") {
		for _, c2 := range strings.Split(second, "") {
			if c1 == c2 {
				return c1
			}
		}
	}

	return ""
}

func getCharPoints(s string) int {
	r := rune(s[0])
	if unicode.IsLower(r) {
		return int(r) - 96
	}

	return int(r) - 38
}

func sumPriorityChars(input [][2]string) (sum int) {
	for _, ruck := range input {
		sum += getCharPoints(findCommonChar(ruck))
	}
	return
}

func findCommonBadge(elfs [][2]string) string {
	for i, elf := range elfs {
		s := elf[0] + elf[1]
		for _, char := range strings.Split(s, "") {
			e1 := elfs[i+1][0] + elfs[i+1][1]
			e2 := elfs[i+2][0] + elfs[i+2][1]
			if strings.Contains(e1, char) && strings.Contains(e2, char) {
				return char
			}
		}
	}
	return ""
}

func groupByThrees(elfs [][2]string) [][][2]string {
	var grouped [][][2]string
	var stack [][2]string
	for i, elf := range elfs {
		stack = append(stack, elf)
		if (i+1)%3 == 0 {
			grouped = append(grouped, stack)
			stack = [][2]string{}
		}
	}

	return grouped
}

func sumBadges(parsed [][2]string) (sum int) {
	elfs := groupByThrees(parsed)
	for _, group := range elfs {
		sum += getCharPoints(findCommonBadge(group))
	}
	return
}
