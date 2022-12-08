package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var example = `30373
25512
65332
33549
35390`

func TestSolutionPart2(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	forest := parse(file)
	assert.Equal(t, 284648, highestScore(forest))
}

func TestGetScenicScoreForCoordinates3And2(t *testing.T) {
	forest := parse(strings.NewReader(example))
	y, x := 3, 2
	assert.Equal(t, 8, GetScenicScore(forest, y, x))
}

func TestGetScenicScoreForCoordinates1And2(t *testing.T) {
	forest := parse(strings.NewReader(example))
	y, x := 1, 2
	assert.Equal(t, 4, GetScenicScore(forest, y, x))
}

func TestSolutionPart1(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	in := parse(file)
	fmt.Println(countVisibleTrees(in))
	assert.Equal(t, 1733, countVisibleTrees(in))
}

func TestCountVisibleTrees(t *testing.T) {
	in := parse(strings.NewReader(example))
	assert.Equal(t, 21, countVisibleTrees(in))
}

func TestParse_parses_to_a_multidimensional_slice_of_int(t *testing.T) {
	expected := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	assert.Equal(t, expected, parse(strings.NewReader(example)))
}
