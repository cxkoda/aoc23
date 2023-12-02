package main

import (
	"fmt"
	"os"
	"strings"

	_ "embed"
)

type Game struct {
	ID      int
	Samples []CubeSet
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func (s *CubeSet) Power() int {
	return s.Red * s.Green * s.Blue
}

func parseGame(s string) (*Game, error) {
	var g Game
	_, err := fmt.Sscanf(s, "Game %d", &g.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing game id in %q: %v", s, err)
	}

	samples := strings.Split(s, ":")[1]
	for _, sample := range strings.Split(samples, ";") {
		var s CubeSet

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

func (g *Game) isPossible() bool {
	for _, s := range g.Samples {
		if s.Red > 12 || s.Green > 13 || s.Blue > 14 {
			return false
		}
	}

	return true
}

func (g *Game) minimumPool() *CubeSet {
	var min CubeSet
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
var input string

func run() error {
	ans1, err := sumPossibleGameIDs(input)
	if err != nil {
		return err
	}
	fmt.Printf("answer 1: %d\n", ans1)

	ans2, err := sumMinimumPoolPowers(input)
	if err != nil {
		return err
	}
	fmt.Printf("answer 2: %d\n", ans2)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
