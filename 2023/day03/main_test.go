package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	matrix := parseInput(`
...*......
..35..633.`)

	res := solve(matrix)

	assert.True(t, isSymbol('$'))
	assert.False(t, isSymbol('.'))
	assert.False(t, isSymbol('.'))

	assert.Equal(t, 35, res.part1)
}
