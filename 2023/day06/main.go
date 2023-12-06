package main

import (
    "fmt"
    _ "embed"
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

func solve(input string) Result {
    var res Result

    // ...

    return res
}
