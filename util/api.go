package util

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var session = "~/.aoc_session"
var client http.Client
var year = time.Now().Year()
var day = time.Now().Day()
var Secret string

// SubmitAns submits the answer to the today's AOC problem.
func SubmitAns(ans string) bool {
	correct, _ := submit(ans, getSubmitUrl(year, day))
	return correct
}

// SubmitAnsForYearDay submits an answer to a previous AOC day.
func SubmitAnsForYearDay(ans string, year, day int) bool {
	correct, _ := submit(ans, getSubmitUrl(year, day))
	return correct
}

func submit(ans, url string) (bool, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(ans))
	if err != nil {
		return false, err
	}
	if Secret == "" {
		Secret = getSecret()
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: Secret})
	resp, err := client.Do(req)
	return resp.StatusCode == 200, err
}

func getSecret() string {
	path, err := ExpandPath(session)
	check(err)
	f, err := os.ReadFile(path)
	check(err)
	return string(f)
}

func ExpandPath(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func getSubmitUrl(year, day int) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
}
