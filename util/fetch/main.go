package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
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
	plainUrl := getUrl(*year, *day)
	req, err := http.NewRequest("GET", plainUrl, nil)
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
	io.Copy(file, resp.Body)
}

func getSecret() {
	path, err := expand(*session)
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

func expand(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}
