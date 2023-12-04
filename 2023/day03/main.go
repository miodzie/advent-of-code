package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input
var input string

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

type Coord struct {
	X int
	Y int
}

func coord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func (c Coord) valid(maxX int, maxY int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < maxX && c.Y < maxY
}

type Gear struct {
	nums []int
}
type Num struct {
	buffer    string
	value     int
	coords    []Coord
	isPartNum bool
}

func (n Num) wasGearNumber(numsAroundGear map[Coord]*Gear) *Gear {
	for _, cord := range n.coords {
		if g, ok := numsAroundGear[cord]; ok {
			return g
		}
	}
	return nil
}

type Result struct {
	part1 int
	part2 int
}

func main() {
	res := solve(parseInput(input))
	fmt.Println(res.part1)
	fmt.Println(res.part2)
}

func solve(matrix [][]rune) Result {
	var res Result
	var num Num
	var gears = make(map[Coord]*Gear)
	var numsAroundGear = make(map[Coord]*Gear)

	for x, line := range matrix {
		for y, c := range line {
			if unicode.IsNumber(c) {
				num.coords = append(num.coords, coord(x, y))
				num.buffer += string(c)

				for _, cord := range allSides(x, y) {
					if !cord.valid(len(matrix), len(line)) {
						continue
					}

					if isSymbol(matrix[cord.X][cord.Y]) {
						num.isPartNum = true
					}

					if matrix[cord.X][cord.Y] == '*' {
						g, ok := gears[cord]
						if !ok {
							g = &Gear{}
							gears[cord] = g
						}
						numsAroundGear[coord(x, y)] = g
					}
				}
				continue
			}

			if num.isPartNum {
				num.value, _ = strconv.Atoi(num.buffer)
				res.part1 += num.value

				if g := num.wasGearNumber(numsAroundGear); g != nil {
					g.nums = append(g.nums, num.value)
					if len(g.nums) == 2 {
						res.part2 += g.nums[0] * g.nums[1]
					}
				}

			}
			num = Num{}
		}
	}

	return res
}

func allSides(x, y int) []Coord {
	return []Coord{
		coord(x-1, y),   // up
		coord(x+1, y),   // down
		coord(x, y-1),   // left
		coord(x, y+1),   // right
		coord(x-1, y+1), // top right
		coord(x-1, y-1), // top left
		coord(x+1, y-1), // bottom left
		coord(x+1, y+1), // bottom right
	}
}

func isSymbol(r rune) bool {
	s := string(r)
	return s != "." && !unicode.IsNumber(r)
}

func parseInput(input string) (matrix [][]rune) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var a []rune
		for _, r := range strings.Split(line, "") {
			a = append(a, rune(r[0]))
		}
		matrix = append(matrix, a)
	}
	return matrix
}
