package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSet(t *testing.T) {
	var set map[string]int
	set = ParseSet("3 blue, 4 red;")
	assert.Equal(t, 3, set["blue"])
	assert.Equal(t, 4, set["red"])
}
