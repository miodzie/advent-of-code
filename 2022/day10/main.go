package main

import (
	"bufio"
	"fmt"
	"os"
)

type operation struct {
	nop    bool
	X      int
	cycles int
}

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	var ops []operation
	for scanner.Scan() {
		line := scanner.Text()
		var op = operation{nop: true}
		if line[0] == 'a' {
			fmt.Sscanf(line, "addx %d", &op.X)
			op.nop = false
		}
		ops = append(ops, op)
	}
	cycles := 0
	X := 1
	var sum, product int

	for _, op := range ops {
		cycles++
		if product = check(cycles, X); product != 0 {
			sum += product
		}
		if !op.nop {
			cycles++
			if product = check(cycles, X); product != 0 {
				sum += product
			}
			X += op.X
		}
	}
	fmt.Println(sum)

	X = 1
	op, ops := ops[0], ops[1:]
	art := ""
	position := 0
	for cycle := 1; cycle < 240; cycle++ {
		if position == 40 {
			position = 0
		}
		if cycle != 1 && (op.nop || op.cycles == 2) {
			op, ops = ops[0], ops[1:]
		}
		op.cycles++
		if X >= position-1 && X <= position+1 {
			art += "#"
		} else {
			art += "."
		}
		art += linebreak(cycle)
		if op.cycles == 2 {
			X += op.X
		}
		position++
	}
	fmt.Print(art)
}

func linebreak(cycle int) string {
	for i := 40; i < 240; i += 40 {
		if cycle == i {
			return "\n"
		}
	}
	return ""
}

func check(cycles int, X int) int {
	for i := 20; i <= 220; i += 40 {
		if cycles == i {
			//fmt.Printf("cycle %d: X: %d\n", cycles, X)
			return X * cycles
		}
	}
	return 0
}
