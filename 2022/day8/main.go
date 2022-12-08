package main

import (
	"bufio"
	. "github.com/miodzie/advent-of-code/util"
	"io"
	"strconv"
	"strings"
)

func highestScore(forest [][]int) int {
	highest := -1

	for y, row := range forest {
		for x := range row {
			treeScore := GetScenicScore(forest, y, x)
			if treeScore > highest {
				highest = treeScore
			}
		}
	}

	return highest
}

func GetScenicScore(forest [][]int, y int, x int) int {
	return score(Reverse(Column(forest, x)[:y]), forest[y][x]) * // up
		score(Reverse(forest[y][:x]), forest[y][x]) * // left
		score(Column(forest, x)[y+1:], forest[y][x]) * // down
		score(forest[y][x+1:], forest[y][x]) // right
}

func score(trees []int, value int) (sum int) {
	for _, t := range trees {
		sum++
		if t >= value {
			break
		}
	}
	return
}

func countVisibleTrees(forest [][]int) (visible int) {
	for y, outer := range forest {
		for x := range outer {
			if isOutside(forest, y, x) {
				visible++
			} else if !isInvisible(forest, y, x) {
				visible++
			}
		}
	}
	return visible
}

func isInvisible(forest [][]int, y int, x int) bool {
	tree := forest[y][x]

	return anyGreaterOrEqualTo(Column(forest, x)[:y], tree) && // up
		anyGreaterOrEqualTo(Column(forest, x)[y+1:], tree) && // below
		anyGreaterOrEqualTo(forest[y][:x], tree) && // left
		anyGreaterOrEqualTo(forest[y][x+1:], tree) // right
}

func anyGreaterOrEqualTo(forest []int, current int) bool {
	for _, tree := range forest {
		if tree >= current {
			return true
		}
	}
	return false
}

func isOutside(forest [][]int, i int, k int) bool {
	if i == 0 || i == len(forest[i])-1 {
		return true
	}
	if (i >= 1 && i <= len(forest[i])-2) &&
		(k == 0 || k == len(forest[i])-1) {
		return true
		//fmt.Printf("%d,%d: %d\n", i, k, forest[i][k])
	}
	return false
}

func parse(r io.Reader) (trees [][]int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "")
		var line []int
		for _, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			line = append(line, i)
		}
		trees = append(trees, line)
	}
	return trees
}
