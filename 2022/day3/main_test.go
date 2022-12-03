package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n"
	expected := [][2]string{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}}

	assert.Equal(t, expected, parseInput(input))
}

func TestFindCommonChar(t *testing.T) {
	given := [2]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}
	expected := "p"

	assert.Equal(t, expected, findCommonChar(given))
}

func TestGetCharPoints(t *testing.T) {
	assert.Equal(t, 1, getCharPoints("a"))
	assert.Equal(t, 27, getCharPoints("A"))
	assert.Equal(t, 16, getCharPoints("p"))
	assert.Equal(t, 38, getCharPoints("L"))
}

func TestSumPriorityChars(t *testing.T) {
	f, err := os.ReadFile("example")
	if err != nil {
		panic(err)
	}
	input := parseInput(string(f))

	assert.Equal(t, 157, sumPriorityChars(input))
}

func TestPart1Solution(t *testing.T) {
	input := getInput()
	assert.Equal(t, 8185, sumPriorityChars(input))
}

func TestFindCommonBadge(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\n"
	elfs := parseInput(input)
	assert.Equal(t, "r", findCommonBadge(elfs))
}

func TestGroupByThrees(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL" +
		"\nPmmdzqPrVvPwwTWBwg\n1234\n"
	expected := [][][2]string{{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
	}}

	assert.Equal(t, expected, groupByThrees(parseInput(input)))
}

func TestSumBadges(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL" +
		"\nPmmdzqPrVvPwwTWBwg\n1234\n"

	assert.Equal(t, 18, sumBadges(parseInput(input)))
}

func TestPart2Solution(t *testing.T) {
	input := getInput()
	assert.Equal(t, 2817, sumBadges(input))
}

func getInput() [][2]string {
	f, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	return parseInput(string(f))
}
