package day01

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/cxkoda/aoc23/days"
)

func ParseDigit(c byte) (int, bool) {
	// Using the fact that '0'..'9' are consecutive in ASCII
	if '0' <= c && c <= '9' {
		return int(c - '0'), true
	}

	return 0, false
}

func ParseDigitPart1(s string, i int) (int, bool) {
	return ParseDigit(s[i])
}

func parseDigitPart2(s string, i int) (int, bool) {
	d, ok := ParseDigitPart1(s, i)
	if ok {
		return d, ok
	}

	for j, word := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		if strings.HasPrefix(s[i:], word) {
			return j, true
		}
	}

	return 0, false
}

type parseDigitFunc func(s string, i int) (int, bool)

func findFirstDigit(s string, parse parseDigitFunc) (int, error) {
	for i := 0; i < len(s); i++ {
		x, done := parse(s, i)
		if !done {
			continue
		}
		return x, nil
	}

	return 0, fmt.Errorf("no digit found in %q", s)
}

func findLastDigit(s string, parse parseDigitFunc) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		x, done := parse(s, i)
		if !done {
			continue
		}
		return x, nil
	}

	return 0, fmt.Errorf("no digit found in %q", s)
}

func decodeCalibration(s string, parse parseDigitFunc) (int, error) {
	a, err := findFirstDigit(s, parse)
	if err != nil {
		return 0, fmt.Errorf("findFirstDigit(%q): %v", s, err)
	}

	b, err := findLastDigit(s, parse)
	if err != nil {
		return 0, fmt.Errorf("findLastDigit(%q): %v", s, err)
	}

	return 10*a + b, nil
}

func totalCalibration(ss []string, parse parseDigitFunc) (int, error) {
	total := 0
	for _, s := range ss {
		n, err := decodeCalibration(s, parse)
		if err != nil {
			return 0, err
		}

		total += n
	}

	return total, nil
}

//go:embed input
var input Input

type Input string

func init() {
	days.MustRegister(1, input)
}

func (i Input) Part1() (int, error) {
	return totalCalibration(strings.Split(string(i), "\n"), ParseDigitPart1)
}

func (i Input) Part2() (int, error) {
	return totalCalibration(strings.Split(string(i), "\n"), parseDigitPart2)
}
