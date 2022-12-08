package main

import (
	"bufio"
	"fmt"
	"github.com/miodzie/advent-of-code/util"
	"io"
	"unicode"
)

type Move struct {
	Amount, From, To int
}
type Crate rune

func ApplyCrateMover9001Moves(moves []Move, stacks []Stack) {
	for _, move := range moves {
		var popped Stack
		for i := 0; i < move.Amount; i++ {
			popped.Push(stacks[move.From-1].Pop())
		}
		for i := len(popped) - 1; i >= 0; i-- {
			stacks[move.To-1].Push(popped[i])
		}
	}
}

func PopOnceAll(stacks []Stack) (crates string) {
	for _, s := range stacks {
		if len(s) > 0 {
			crates += string(s.Pop())
		}
	}
	return crates
}

func ApplyMoves(moves []Move, stacks []Stack) {
	for _, m := range moves {
		stacks = ApplyMove(m, stacks)
	}
}

func ApplyMove(move Move, stacks []Stack) []Stack {
	for i := 0; i < move.Amount; i++ {
		stacks[move.To-1].Push(stacks[move.From-1].Pop())
	}
	return stacks
}

type Stack []Crate

func (s *Stack) Push(crate Crate) {
	*s = append(*s, crate)
}

func (s *Stack) Prepend(crate Crate) {
	*s = util.Prepend(crate, *s)
}

func (s *Stack) Pop() Crate {
	stack := *s
	popped := stack[len(stack)-1]
	*s = stack[:len(stack)-1]
	return popped
}

func parseInput(reader io.Reader) ([]Stack, []Move) {
	scanner := bufio.NewScanner(reader)
	var moves []Move
	var stacks = []Stack{{}}
	for scanner.Scan() {
		// Crates
		line := []rune(scanner.Text())
		k := 0
		for j := 1; j < len(line); j += 4 {
			if unicode.IsUpper(line[j]) {
				if len(stacks) < k+1 {
					for i := len(stacks); i < k+1; i++ {
						stacks = append(stacks, Stack{})
					}
				}
				stacks[k].Prepend(Crate(line[j]))
			}
			k++
		}
		// Moves
		var move Move
		fmt.Sscanf(scanner.Text(),
			"move %d from %d to %d",
			&move.Amount, &move.From, &move.To)
		if move.Amount != 0 {
			moves = append(moves, move)
		}
	}

	return stacks, moves
}
