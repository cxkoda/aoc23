package day02

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/cxkoda/aoc23/days"
)

type game struct {
	ID      int
	Samples []cubeSet
}

type cubeSet struct {
	Red   int
	Green int
	Blue  int
}

func (s *cubeSet) Power() int {
	return s.Red * s.Green * s.Blue
}

func parseGame(s string) (*game, error) {
	var g game
	_, err := fmt.Sscanf(s, "Game %d", &g.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing game id in %q: %v", s, err)
	}

	samples := strings.Split(s, ":")[1]
	for _, sample := range strings.Split(samples, ";") {
		var s cubeSet

		for _, cubes := range strings.Split(sample, ",") {
			var color string
			var num int
			if _, err := fmt.Sscanf(cubes, "%d %s", &num, &color); err != nil {
				return nil, fmt.Errorf("parsing number of %s in %q: %v", color, sample, err)
			}

			switch color {
			case "red":
				s.Red = num
			case "green":
				s.Green = num
			case "blue":
				s.Blue = num
			}
		}

		g.Samples = append(g.Samples, s)
	}

	return &g, nil
}

func (g *game) isPossible() bool {
	for _, s := range g.Samples {
		if s.Red > 12 || s.Green > 13 || s.Blue > 14 {
			return false
		}
	}

	return true
}

func (g *game) minimumPool() *cubeSet {
	var min cubeSet
	for _, s := range g.Samples {
		if s.Red > min.Red {
			min.Red = s.Red
		}

		if s.Green > min.Green {
			min.Green = s.Green
		}

		if s.Blue > min.Blue {
			min.Blue = s.Blue
		}
	}

	return &min
}

func sumPossibleGameIDs(games string) (int, error) {
	total := 0
	for _, s := range strings.Split(games, "\n") {
		g, err := parseGame(s)
		if err != nil {
			return 0, err
		}

		if g.isPossible() {
			total += g.ID
		}
	}

	return total, nil
}

func sumMinimumPoolPowers(games string) (int, error) {
	total := 0
	for _, s := range strings.Split(games, "\n") {
		g, err := parseGame(s)
		if err != nil {
			return 0, err
		}

		total += g.minimumPool().Power()
	}

	return total, nil
}

//go:embed input
var input Input

type Input string

func init() {
	days.MustRegister(2, input)
}

func (in Input) Part1() (int, error) {
	return sumPossibleGameIDs(string(in))
}

func (in Input) Part2() (int, error) {
	return sumMinimumPoolPowers(string(in))
}
