package template

import (
	_ "embed"

	"github.com/cxkoda/aoc23/days"
)

//go:embed input
var input Challenge

type Challenge string

func init() {
	days.MustRegister(0, input)
}

func (in Challenge) Part1() (int, error) {
	return 0, nil
}

func (in Challenge) Part2() (int, error) {
	return 0, nil
}
