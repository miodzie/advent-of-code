package util

func Column[T comparable](grid [][]T, x int) (col []T) {
	for _, row := range grid {
		col = append(col, row[x])
	}
	return
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

func Prepend[T any](item T, items []T) []T {
	return append([]T{item}, items[:]...)
}

func Reverse[T any](slice []T) (reversed []T) {
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}
	return
}

func Abs[T float64 | int | int32](n T) T {
	if n > 0 {
		return n
	}
	return -n
}
