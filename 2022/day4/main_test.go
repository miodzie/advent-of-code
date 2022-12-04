package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var sample = []Pair{
	{Section{2, 4}, Section{6, 8}},
	{Section{2, 3}, Section{4, 5}},
	{Section{5, 7}, Section{7, 9}},
	{Section{2, 8}, Section{3, 7}},
	{Section{6, 6}, Section{4, 6}},
	{Section{2, 6}, Section{4, 8}},
}

func TestParseInput(t *testing.T) {
	f, err := os.Open("example")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, sample, parseInput(f))
}

func TestExampleSolutions(t *testing.T) {
	f, err := os.Open("example")
	if err != nil {
		panic(err)
	}
	pairs := parseInput(f)
	assert.Equal(t, 2, fullyContainedPairs(pairs))
	assert.Equal(t, 4, overlappingPairs(pairs))
}

func TestSolutions(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	pairs := parseInput(f)
	assert.Equal(t, 644, fullyContainedPairs(pairs))
	assert.Equal(t, 926, overlappingPairs(pairs))
}

func TestSection_Contains_another_section(t *testing.T) {
	s1 := Section{2, 8}
	s2 := Section{3, 7}

	assert.True(t, s1.Contains(s2))
}

func TestSection_Contains_does_not_contain_another_section(t *testing.T) {
	s1 := Section{2, 8}
	s2 := Section{3, 7}

	assert.False(t, s2.Contains(s1))
}

func TestSection_Overlap_is_true_when_other_section_overlaps(t *testing.T) {
	s1 := Section{5, 7}
	s2 := Section{7, 9}
	assert.True(t, s1.Overlaps(s2))
}

func TestSection_Overlap_is_false_when_other_section_doesnt_overlap(t *testing.T) {
	s1 := Section{2, 4}
	s2 := Section{6, 8}
	assert.False(t, s1.Overlaps(s2))
	assert.False(t, s2.Overlaps(s1))
}

func Test_CountFullyContainedPairs(t *testing.T) {
	assert.Equal(t, 2, fullyContainedPairs(sample))
}
