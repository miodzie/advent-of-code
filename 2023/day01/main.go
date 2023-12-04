package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input
var input string

var example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func main() {
	part1(input)
	part2(input)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	var sum int
	key := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var keys []string
	for k, _ := range key {
		keys = append(keys, k)
	}

	for _, line := range lines {
		var first, last string
		var buff string
		for _, c := range line {
			buff += string(c)
			if k, ok := matchesKey(buff, keys); ok {
				first = key[k]
				buff = ""
				break
			}
			if unicode.IsNumber(c) {
				first = string(c)
				buff = ""
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			buff = string(c) + buff

			if k, ok := matchesKey(buff, keys); ok {
				last = key[k]
				buff = ""
				break
			}
			if unicode.IsNumber(rune(c)) {
				last = string(c)
				buff = ""
				break
			}
		}

		num, _ := strconv.Atoi(first + last)

		sum += num
	}

	fmt.Println("part2: ", sum)
}

func matchesKey(buff string, keys []string) (string, bool) {
	for _, k := range keys {
		if strings.Contains(buff, k) {
			return k, true
		}
	}
	return "", false
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	var sums int
	for _, line := range lines {
		// get first
		var first, last string
		for _, c := range line {
			if unicode.IsNumber(c) {
				first = string(c)
				break
			}
		}

		// get last
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				last = string(line[i])
				break
			}
		}

		num, _ := strconv.Atoi(first + last)
		// sum
		sums += num
	}
	fmt.Println("part1: ", sums)
}
