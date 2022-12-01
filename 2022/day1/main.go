package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input")
	elfs := strings.Split(string(f), "\n\n")
	sums := []int{0, 0, 0}
	for _, calories := range elfs {
		sum := 0
		for _, c := range strings.Split(calories, "\n") {
			ci, _ := strconv.Atoi(c)
			sum += ci
		}
		sums = append(sums, sum)
		sort.Sort(sort.Reverse(sort.IntSlice(sums)))
		sums = sums[:3]
	}
	fmt.Println(sums[0])
	fmt.Println(sums[0] + sums[1] + sums[2])
}
