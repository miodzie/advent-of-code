package main

import "strings"

func findMarker(packet string, n int) int {
	letters := strings.Split(packet, "")
	for i := 0; i < len(letters); i++ {
		section := letters[i : i+n]
		if IsUnique(section) {
			return i + n
		}
	}

	return -1
}

func IsUnique[T comparable](items []T) bool {
	seen := make(map[T]any)
	for _, e := range items {
		_, exists := seen[e]
		if exists {
			return false
		}
		seen[e] = nil
	}
	return true
}
