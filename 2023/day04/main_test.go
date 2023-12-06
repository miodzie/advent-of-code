package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	res := solve(example)
	assert.Equal(t, 13, res.part1)
	assert.Equal(t, 30, res.part2)
}


func TestSolve(t *testing.T) {
	res := solve(input)
	assert.Equal(t, res.part1, 17782)
	assert.Equal(t, res.part2, 8477787)
}

func Test_parseNums(t *testing.T) {
	nums, _ := parseNums("14  1")

	fmt.Println(nums)
}
