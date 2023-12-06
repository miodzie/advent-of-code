package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	res := solve(example)

	assert.Equal(t, 288, res.part1)
	assert.Equal(t, 71503, res.part2)
}

func TestSolve(t *testing.T) {
	res := solve(input)

	assert.Equal(t, 781200, res.part1)
	assert.Equal(t, 49240091, res.part2)
}

