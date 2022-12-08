package main

import (
	"github.com/miodzie/advent-of-code/util"
	"strings"
)

func findMarker(packet string, n int) int {
	letters := strings.Split(packet, "")
	for i := 0; i < len(letters); i++ {
		section := letters[i : i+n]
		if util.IsUnique(section) {
			return i + n
		}
	}

	return -1
}
