package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseGame(t *testing.T) {
	tests := []struct {
		input string
		want  *Game
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			&Game{
				ID: 1,
				Samples: []CubeSet{
					{Red: 4, Green: 0, Blue: 3},
					{Red: 1, Green: 2, Blue: 6},
					{Red: 0, Green: 2, Blue: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		got, err := parseGame(tt.input)
		if err != nil {
			t.Fatalf("parseGame(%q) error %v, want %v", tt.input, err, nil)
		}

		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Errorf("parseGame(%q) diff(+got -want) %v", tt.input, diff)
		}
	}
}

func TestGamePossible(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			true,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			true,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			false,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			false,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			true,
		},
	}
	for _, tt := range tests {
		got, err := parseGame(tt.input)
		if err != nil {
			t.Fatalf("parseGame(%q) error %v, want %v", tt.input, err, nil)
		}

		if got.isPossible() != tt.want {
			t.Errorf("parseGame(%q).isPossible() = %v, want %v", tt.input, got.isPossible(), tt.want)
		}
	}
}

func TestSumPossibleGameIDs(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			8,
		},
	}
	for _, tt := range tests {
		got, err := sumPossibleGameIDs(tt.input)
		if err != nil {
			t.Fatalf("sumPossibleGameIDs(%q) error %v, want %v", tt.input, err, nil)
		}

		if got != tt.want {
			t.Errorf("sumPossibleGameIDs(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestMinimumPool(t *testing.T) {
	tests := []struct {
		input string
		want  CubeSet
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			CubeSet{Red: 4, Green: 2, Blue: 6},
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			CubeSet{Red: 1, Green: 3, Blue: 4},
		},
	}
	for _, tt := range tests {
		got, err := parseGame(tt.input)
		if err != nil {
			t.Fatalf("parseGame(%q) error %v, want %v", tt.input, err, nil)
		}

		if diff := cmp.Diff(tt.want, *got.minimumPool()); diff != "" {
			t.Errorf("parseGame(%q).minimumPool() diff(+got -want) %v", tt.input, diff)
		}
	}
}

func TestMinimumPoolPower(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			48,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			12,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			1560,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			630,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			36,
		},
	}
	for _, tt := range tests {
		game, err := parseGame(tt.input)
		if err != nil {
			t.Fatalf("parseGame(%q) error %v, want %v", tt.input, err, nil)
		}

		got := game.minimumPool().Power()
		if got != tt.want {
			t.Errorf("%+v.minimumPool().Power() = %d, want %v", game, got, tt.want)
		}
	}
}

func TestSumMinimumPoolPowers(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			2286,
		},
	}
	for _, tt := range tests {
		got, err := sumMinimumPoolPowers(tt.input)
		if err != nil {
			t.Fatalf("sumMinimumPoolPowers(%q) error %v, want %v", tt.input, err, nil)
		}

		if got != tt.want {
			t.Errorf("sumMinimumPoolPowers(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
