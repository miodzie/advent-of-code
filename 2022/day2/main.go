package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := parseInput(string(f))
	fmt.Println(sumScores(input))
	sum := 0
	for _, round := range input {
		sum += calculateScore(applyStrategy(round))
	}
	fmt.Println(sum)
}

const (
	LOSE = "X"
	DRAW = "Y"
	WIN  = "Z"
)

func applyStrategy(round [2]string) [2]string {
	opponent := round[0]
	strategy := round[1]
	if strategy == DRAW {
		round[1] = opponent
	}
	if strategy == LOSE {
		if opponent == _PAPER {
			round[1] = _ROCK
		}
		if opponent == _SCISSORS {
			round[1] = _PAPER
		}
		if opponent == _ROCK {
			round[1] = _SCISSORS
		}
	}
	if strategy == WIN {
		if opponent == _ROCK {
			round[1] = _PAPER
		}
		if opponent == _PAPER {
			round[1] = _SCISSORS
		}
		if opponent == _SCISSORS {
			round[1] = _ROCK
		}
	}

	return round
}

const (
	ROCK      = 1
	PAPER     = 2
	SCISSORS  = 3
	_ROCK     = "A"
	_PAPER    = "B"
	_SCISSORS = "C"
)

var legend = map[string]int{
	"A": ROCK, "B": PAPER, "C": SCISSORS,
	"X": ROCK, "Y": PAPER, "Z": SCISSORS,
}

func calculateScore(round [2]string) int {
	opponent, _ := legend[round[0]]
	me, _ := legend[round[1]]
	var score int
	// Tie
	if opponent == me {
		return me + 3
	}
	// Wins
	if me == ROCK && opponent == SCISSORS ||
		me == PAPER && opponent == ROCK ||
		me == SCISSORS && opponent == PAPER {
		score += 6
	}
	// If we didn't win, at least get the hand we played.
	score += me

	return score
}

func sumScores(rounds [][2]string) (sum int) {
	for _, r := range rounds {
		sum += calculateScore(r)
	}
	return
}

func parseInput(input string) [][2]string {
	var results [][2]string
	split := strings.Split(input, "\n")
	for _, s := range split {
		chars := strings.Split(s, " ")
		if len(chars) > 1 {
			results = append(results, [2]string{chars[0], chars[1]})
		}
	}

	return results
}
