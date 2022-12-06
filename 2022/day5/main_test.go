package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var sample = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

var exampleInput = strings.NewReader(sample)

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	// TODO: Parse stacks automatically guh.
	stacks, moves := parseInput(f)
	ApplyMoves(moves, stacks)
	assert.Equal(t, "WSFTMRHPP", PopOnceAll(stacks))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	// TODO: Parse stacks automatically guh.
	stacks, moves := parseInput(f)
	ApplyCrateMover9001Moves(moves, stacks)
	assert.Equal(t, "GSLCMFBRP", PopOnceAll(stacks))
}

func TestExampleSolutions(t *testing.T) {
	// Part 1
	stacks := []Stack{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}
	moves := []Move{
		{Amount: 1, From: 2, To: 1},
		{Amount: 3, From: 1, To: 3},
		{Amount: 2, From: 2, To: 1},
		{Amount: 1, From: 1, To: 2},
	}
	ApplyMoves(moves, stacks)
	assert.Equal(t, "CMZ", PopOnceAll(stacks))

	// Part 2
	stacks = []Stack{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}
	ApplyCrateMover9001Moves(moves, stacks)
	assert.Equal(t, "MCD", PopOnceAll(stacks))
}

func TestApplyMove(t *testing.T) {
	stacks := []Stack{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}
	move := Move{Amount: 1, From: 2, To: 1}
	expected := []Stack{{'Z', 'N', 'D'}, {'M', 'C'}, {'P'}}

	assert.Equal(t, expected, ApplyMove(move, stacks))
}

func TestApplyCrateMover9001Moves(t *testing.T) {
	stacks := []Stack{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}
	moves := []Move{{Amount: 1, From: 2, To: 1}, {Amount: 3, From: 1, To: 3}}
	expected := []Stack{{}, {'M', 'C'}, {'P', 'Z', 'N', 'D'}}
	ApplyCrateMover9001Moves(moves, stacks)
	assert.Equal(t, expected, stacks)
}

func TestApplyCrateMover9001Moves2(t *testing.T) {
	stacks := []Stack{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}
	moves := []Move{{Amount: 1, From: 2, To: 1},
		{Amount: 3, From: 1, To: 3},
		{Amount: 2, From: 2, To: 1}}
	expected := []Stack{{'M', 'C'}, {}, {'P', 'Z', 'N', 'D'}}
	ApplyCrateMover9001Moves(moves, stacks)
	assert.Equal(t, expected, stacks)
}

func TestParseInput(t *testing.T) {
	stacks, moves := parseInput(exampleInput)
	expected := []Stack{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	expectedMoves := []Move{
		{Amount: 1, From: 2, To: 1},
		{Amount: 3, From: 1, To: 3},
		{Amount: 2, From: 2, To: 1},
		{Amount: 1, From: 1, To: 2},
	}

	assert.Equal(t, expected, stacks)
	assert.Equal(t, expectedMoves, moves)
	assert.Len(t, stacks, 3)
}
