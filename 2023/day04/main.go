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

var example = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

type Result struct {
	part1 int
	part2 int
}

func main() {
	res := solve(input)
	fmt.Println(res.part1)
	fmt.Println(res.part2)
}

type Card struct {
	Number  int
	Winners map[int]int
	Numbers []int
	Wins    int
}

func solve(input string) Result {
	var res Result
	var cards = []*Card{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		pointValue := 0
		card := parseCard(line)

		cards = append(cards, card)

		for _, num := range card.Numbers {
			if _, ok := card.Winners[num]; ok {
				card.Wins += 1
				pointValue *= 2
				if pointValue == 0 {
					pointValue = 1
				}
			}
		}

		res.part1 += pointValue
	}

	count := map[int]int{}
	for num, card := range cards {
		count[num] += 1
		for i := 1; i <= count[num]; i++ {
			for i := 1; i <= card.Wins; i++ {
				count[num+i]++
			}
		}

	}
	for _, n := range count {
		res.part2 += n
	}

	return res
}

func parseCard(line string) *Card {
	card := &Card{}

	gameStr, cardStr, _ := strings.Cut(line, ":")
	card.Number, _ = strconv.Atoi(string(gameStr[len(gameStr)-1]))

	winS, cardS, _ := strings.Cut(cardStr, "|")
	_, card.Winners = parseNums(winS)
	card.Numbers, _ = parseNums(cardS)

	return card
}

func parseNums(numbers string) ([]int, map[int]int) {
	var nums []int
	var buff string
	var numMap = make(map[int]int)

	numbers += " "
	for _, char := range numbers {
		if unicode.IsNumber(char) {
			buff += string(char)
			continue
		}

		if buff != "" {
			num, _ := strconv.Atoi(buff)
			nums = append(nums, num)
			numMap[num] = 0
			buff = ""
		}
	}

	return nums, numMap
}
