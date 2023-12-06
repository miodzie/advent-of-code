package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

var example = ``

type Result struct {
	part1 int
	part2 int
}

func main() {
	res := solve(example)

	fmt.Println(res.part1)
	fmt.Println(res.part2)
}

func solve(input string) (res Result) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = line
		// ...
	}

	return res
}
