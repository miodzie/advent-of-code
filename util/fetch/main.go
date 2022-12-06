package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

	jar, err := cookiejar.New(&cookiejar.Options{})
	check(err)
	client = http.Client{Jar: jar}
	auth := &http.Cookie{Name: "session", Value: secret}

	plainUrl := getUrl(*year, *day)
	durl, err := url.Parse(plainUrl)
	check(err)
	client.Jar.SetCookies(durl, []*http.Cookie{auth})

	resp, err := client.Get(plainUrl)
	defer resp.Body.Close()
	check(err)

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
