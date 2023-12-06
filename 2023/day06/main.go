package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/miodzie/advent-of-code/parse"
)

//go:embed input
var input string

var example = `Time:      7  15   30
Distance:  9  40  200`

type Result struct {
	part1 int
	part2 int
}

func main() {
	res := solve(input)

	fmt.Println(res.part1)
	fmt.Println(res.part2)
}

type Race struct {
	Time         int
	BestDistance int
}

func (r *Race) WaysToWin() (ways int) {
	for buttonHeld := 1; buttonHeld <= r.Time; buttonHeld++ {
		timeLeft := r.Time - buttonHeld
		distance := timeLeft * buttonHeld

		if distance > r.BestDistance {
			ways++
		}
	}

	return ways
}

func solve(input string) (res Result) {
	races, race := parseRaces(input)
	res.part1 = 1

	for _, race := range races {
		res.part1 *= race.WaysToWin()
	}
	res.part2 = race.WaysToWin()

	return res
}

func parseRaces(input string) ([]Race, Race) {
	lines := strings.Split(input, "\n")
	times := parse.Nums(lines[0])
	distances := parse.Nums(lines[1])

	var races []Race
	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			Time:         times[i],
			BestDistance: distances[i],
		})
	}

	return races, Race{
		Time:         parse.Nums(strings.ReplaceAll(lines[0], " ", ""))[0],
		BestDistance: parse.Nums(strings.ReplaceAll(lines[1], " ", ""))[0],
	}
}
