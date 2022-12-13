package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("example")
	var packets [][]any
	for _, pair := range strings.Split(string(f), "\n\n") {
		s := strings.Split(pair, "\n")
		var a, b []int
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		packets = append(packets, a)
		packets = append(packets, b)
	}
	fmt.Println(packets)
	fmt.Println(packets[2])
}
