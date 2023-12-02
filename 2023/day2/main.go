package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

var example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func main() {
	part1(input)
	part2(input)
}

func part2(input string) {
	lines := strings.Split(input, "\n")

	var sum int
	for _, line := range lines {
		split := strings.Split(line, ":")
		sets := strings.Split(split[1], ";")

		maxSeen := map[string]int{"green": 1, "blue": 1, "red": 1}
		for _, s := range sets {
			set := ParseSet(s)
			for color, num := range set {
				if _, ok := maxSeen[color]; ok && num > maxSeen[color] {
					maxSeen[color] = num
				}
			}
		}
		sum += maxSeen["green"] * maxSeen["blue"] * maxSeen["red"]
	}
	fmt.Println(sum)
}

func part1(input string) {
	lines := strings.Split(input, "\n")

	var sum int
	for _, line := range lines {
		split := strings.Split(line, ":")
		sets := strings.Split(split[1], ";")

		if setsOk(sets) {
			var game int
			fmt.Sscanf(split[0], "Game %d", &game)
			sum += game
		}
	}
	fmt.Println(sum)
}

func setsOk(sets []string) bool {
	for _, s := range sets {
		set := ParseSet(s)
		if set["red"] > 12 || set["green"] > 13 || set["blue"] > 14 {
			return false
		}
	}
	return true
}

func ParseSet(s string) (set map[string]int) {
	set = make(map[string]int)
	cubes := strings.Split(s, ",")
	for _, cube := range cubes {
		for _, color := range []string{"blue", "red", "green"} {
			if num, found := attemptScan(cube, color); found {
				set[color] = num
			}
		}
	}
	return
}

func attemptScan(cube, color string) (int, bool) {
	var tmp int
	_, err := fmt.Sscanf(cube, "%d "+color, &tmp)
	if err == nil {
		return tmp, true
	}
	return 0, false
}
