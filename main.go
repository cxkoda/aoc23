package main

import (
	"fmt"
	"os"

	"github.com/cxkoda/aoc23/days"

	_ "github.com/cxkoda/aoc23/days/day01"
	_ "github.com/cxkoda/aoc23/days/day02"
	_ "github.com/cxkoda/aoc23/days/day03"
	_ "github.com/cxkoda/aoc23/days/day13"
	_ "github.com/cxkoda/aoc23/days/day15"
)

func main() {
	if err := days.EvaluateAll(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
