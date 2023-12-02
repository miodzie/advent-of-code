package util

import (
	"fmt"
	"golang.org/x/term"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var session = "~/.aoc_session"
var Secret = getSecret()
var client http.Client
var year = time.Now().Year()
var today = time.Now().Day()

func PromptSubmit(ans string) {
	fmt.Printf("Do you want to submit: %s ? [y\\n]", ans)
	yes := yesNoEnter()
	fmt.Print("\n")
	if yes {
		if SubmitAns(ans) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Wrong!")
		}
	}
}

// SubmitAns submits the answer to the today's AOC problem.
// FIXME
func SubmitAns(ans string) bool {
	return SubmitAnsForYearDay(ans, year, today)
}

// SubmitAnsForYearDay submits an answer to a previous AOC today.
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
func yes() bool {
	var y rune
	fmt.Scanf("%c", &y)
	return y == 'y' || y == 'Y'
}

func yesNoEnter() bool {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	check(err)
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	check(err)
	y := b[0]
	return y == 'y' || y == 'Y'
}
