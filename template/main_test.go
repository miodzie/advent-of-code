package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	res := solve(example)

	assert.Equal(t, 0, res.part1)
	// assert.Equal(t, 0, res.part2)
}

func TestSolve(t *testing.T) {
	res := solve(input)

	assert.Equal(t, 0, res.part1)
	// assert.Equal(t, 0, res.part2)
}
