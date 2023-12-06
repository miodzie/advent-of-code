package parse

import (
	"strconv"
	"unicode"
)

func Nums(line string) (nums []int) {
	var buff string

	line += " "
	for _, char := range line {
		if unicode.IsNumber(char) {
			buff += string(char)
			continue
		}

		if buff != "" {
			num, _ := strconv.Atoi(buff)
			nums = append(nums, num)
			buff = ""
		}
	}

	return nums
}
