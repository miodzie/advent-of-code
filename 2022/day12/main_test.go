package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestParseInput(t *testing.T) {
	expected := [][]Node{{}}
	for _, c := range "Sabqponm" {
		expected[0] = append(expected[0], Node{c})
	}
	input := ParseInput(strings.NewReader(example))
	assert.Equal(t, expected[0], input[0])
}
