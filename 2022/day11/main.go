package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id           int
	items        []int
	operation    string
	divisibleBy  int
	throwToTrue  int
	throwToFalse int
	inspections  int
}

func (m *Monkey) operate(item int) int {
	m.inspections++
	if strings.Contains(m.operation, "*") {
		var multiplier int
		var err error
		if strings.Count(m.operation, "old") == 2 {
			return item * item
		} else {
			_, err = fmt.Sscanf(
				m.operation, " Operation: new = old * %d", &multiplier)
		}
		if err != nil {
			fmt.Println(m.operation)
			panic(err)
		}
		return item * multiplier
	}
	if strings.Contains(m.operation, "+") {
		var add int
		_, err := fmt.Sscanf(
			m.operation, " Operation: new = old + %d", &add)
		if err != nil {
			panic(err)
		}
		return item + add
	}
	fmt.Println(m.operation)
	panic("unimplemented")
	return -1
}

func (m *Monkey) test(item int) bool {
	return item%m.divisibleBy == 0
}

func (m *Monkey) catch(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) itemsString() []string {
	var items []string
	for _, i := range m.items {
		items = append(items, strconv.Itoa(i))
	}
	return items
}

func main() {
	monkies := parseInput()
	lcm := 1
	for _, m := range monkies {
		lcm *= m.divisibleBy
	}
	//fmt.Println(lcm)
	for part := 1; part <= 2; part++ {
		monkies = parseInput()
		inspections := make([]int, len(monkies))
		rounds := []int{20, 10_000}
		for round := 0; round < rounds[part-1]; round++ {
			for _, monkey := range monkies {
				for range monkey.items {
					inspections[monkey.id]++
					var item int
					item, monkey.items = monkey.items[0], monkey.items[1:]
					item = monkey.operate(item)
					if part == 1 {
						item = item / 3
					} else {
						item %= lcm
					}
					if monkey.test(item) {
						monkies[monkey.throwToTrue].catch(item)
					} else {
						monkies[monkey.throwToFalse].catch(item)
					}
				}
			}
			//if round == 20-1 {
			//	debug(monkies)
			//}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
		fmt.Println(inspections[0] * inspections[1])
	}
}

func debug(monkies []*Monkey) {
	for _, monkey := range monkies {
		fmt.Printf("Monkey %d, Inspections: %d\n", monkey.id, monkey.inspections)
		//fmt.Printf("Monkey %d: %s\n", monkey.id, strings.Join(monkey.itemsString(), ", "))
	}
}

func parseInput() (monkies []*Monkey) {
	f, _ := os.ReadFile("input")
	_monkies := strings.Split(string(f), "\n")
	for i := 0; i < len(_monkies); i += 7 {
		monkey := parseMonkey(_monkies[i : i+6])
		monkies = append(monkies, &monkey)
	}
	return monkies
}

func parseMonkey(lines []string) Monkey {
	var monkey Monkey
	fmt.Sscanf(lines[0], "Monkey %d:", &monkey.id)
	_items := strings.Split(
		strings.ReplaceAll(lines[1], "Starting items:", ""),
		",")
	for _, item := range _items {
		var i int
		fmt.Sscanf(item, "%d", &i)
		monkey.items = append(monkey.items, i)
	}
	monkey.operation = lines[2]
	fmt.Sscanf(lines[3], " Test: divisible by %d", &monkey.divisibleBy)
	fmt.Sscanf(lines[4], " If true: throw to monkey %d", &monkey.throwToTrue)
	fmt.Sscanf(lines[5], " If false: throw to monkey %d", &monkey.throwToFalse)
	return monkey
}
