package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/miodzie/advent-of-code/util"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var year = flag.Int("year", time.Now().Year(), "the year, defaults to today")
var day = flag.Int("day", time.Now().Day(), "the day, defaults today")
var session = flag.String("secret", "~/.aoc_session", "path to file with session key")
var output = flag.String("out", "input", "file path to write to")

var secret string

var client http.Client

func main() {
	flag.Parse()
	getSecret()
	req, err := http.NewRequest("GET", getUrl(*year, *day), nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: secret})
	check(err)

	// Fetch
	resp, err := client.Do(req)
	defer resp.Body.Close()
	check(err)

	// Write
	file, err := os.Create(*output)
	defer file.Close()
	check(err)
	b, _ := io.ReadAll(resp.Body)
	cleaned := bytes.NewBufferString(strings.TrimSpace(string(b)))
	io.Copy(file, cleaned)
}

func getSecret() {
	path, err := util.ExpandPath(*session)
	check(err)
	f, err := os.ReadFile(path)
	check(err)
	secret = string(f)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getUrl(year, day int) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
}
