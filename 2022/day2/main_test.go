package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrawStrategy(t *testing.T) {
	rounds := [][2]string{{_ROCK, DRAW}, {_PAPER, DRAW}, {_SCISSORS, DRAW}}
	for _, round := range rounds {
		expected := [2]string{round[0], round[0]}
		assert.Equal(t, expected, applyStrategy(round))
	}
}

func TestLoseStrategy(t *testing.T) {
	// Rock
	round := [2]string{_ROCK, LOSE}
	assert.Equal(t, [2]string{_ROCK, _SCISSORS}, applyStrategy(round))
	// Paper
	round = [2]string{_PAPER, LOSE}
	assert.Equal(t, [2]string{_PAPER, _ROCK}, applyStrategy(round))
	// Scissors
	round = [2]string{_SCISSORS, LOSE}
	assert.Equal(t, [2]string{_SCISSORS, _PAPER}, applyStrategy(round))
}

func TestWinStrategy(t *testing.T) {
	// Rock
	round := [2]string{_ROCK, WIN}
	assert.Equal(t, [2]string{_ROCK, _PAPER}, applyStrategy(round))
	// Paper
	round = [2]string{_PAPER, WIN}
	assert.Equal(t, [2]string{_PAPER, _SCISSORS}, applyStrategy(round))
	// Scissors
	round = [2]string{_SCISSORS, WIN}
	assert.Equal(t, [2]string{_SCISSORS, _ROCK}, applyStrategy(round))
}

func TestRoundOne(t *testing.T) {
	round := [2]string{"A", "Y"}
	assert.Equal(t, 8, calculateScore(round))
}

func TestRoundTwo(t *testing.T) {
	round := [2]string{"B", "X"}
	assert.Equal(t, 1, calculateScore(round))
}

func TestRoundThree(t *testing.T) {
	round := [2]string{"C", "Z"}
	assert.Equal(t, 6, calculateScore(round))
}

func TestSumScores(t *testing.T) {
	rounds := [][2]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}
	assert.Equal(t, 15, sumScores(rounds))
}

func TestItParsesInput(t *testing.T) {
	input := "B Y\nA X\nB Y\n"
	expected := [][2]string{{"B", "Y"}, {"A", "X"}, {"B", "Y"}}

	turns := parseInput(input)

	assert.Equal(t, expected, turns)
}
