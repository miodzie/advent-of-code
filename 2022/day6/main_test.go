package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var input, err = os.ReadFile("input")

func TestSolutions(t *testing.T) {
	if err != nil {
		panic(err)
	}
	// Part 1
	assert.Equal(t, 1175, findMarker(string(input), 4))
	// Part 2
	assert.Equal(t, 3217, findMarker(string(input), 14))
}

func TestFindMarker(t *testing.T) {
	packet := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	assert.Equal(t, 5, findMarker(packet, 4))
}

func TestFindMarker14DistinctCharacters(t *testing.T) {
	packet := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	assert.Equal(t, 19, findMarker(packet, 14))
}
