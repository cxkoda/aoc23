package main

import (
	"fmt"
	"os"
	"strings"

	_ "embed"
)

func parseDigit(s string, i int) (int, bool) {
	// Using the fact that '0'..'9' are consecutive in ASCII
	if c := s[i]; '0' <= c && c <= '9' {
		return int(c - '0'), true
	}

	for j, word := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		if strings.HasPrefix(s[i:], word) {
			return j, true
		}
	}

	return 0, false
}

func findFirstDigit(s string) (int, error) {
	for i := 0; i < len(s); i++ {
		x, done := parseDigit(s, i)
		if !done {
			continue
		}
		return x, nil
	}

	return 0, fmt.Errorf("no digit found in %q", s)
}

func findLastDigit(s string) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		x, done := parseDigit(s, i)
		if !done {
			continue
		}
		return x, nil
	}

	return 0, fmt.Errorf("no digit found in %q", s)
}

func decodeCalibration(s string) (int, error) {
	a, err := findFirstDigit(s)
	if err != nil {
		return 0, fmt.Errorf("findFirstDigit(%q): %v", s, err)
	}

	b, err := findLastDigit(s)
	if err != nil {
		return 0, fmt.Errorf("findLastDigit(%q): %v", s, err)
	}

	return 10*a + b, nil
}

func totalCalibration(ss []string) (int, error) {
	total := 0
	for _, s := range ss {
		n, err := decodeCalibration(s)
		if err != nil {
			return 0, err
		}

		total += n
	}

	return total, nil
}

//go:embed input
var input string

func run() error {
	cal, err := totalCalibration(strings.Split(input, "\n"))
	if err != nil {
		return err
	}

	fmt.Println(cal)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
