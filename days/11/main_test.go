package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseUniverse(t *testing.T) {
	tests := []struct {
		input string
		want  *Universe
	}{
		{
			`..#.
....
.#..`,
			&Universe{
				Galaxies: []*Galaxy{
					{2, 0},
					{1, 2},
				},
				DistanceX: []int{2, 1, 1, 2},
				DistanceY: []int{1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		got, err := parseUniverse(tt.input, 2)
		if err != nil {
			t.Fatalf("parseUniverse(%q) error %v, want %v", tt.input, err, nil)
		}

		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Errorf("parseUniverse(%q) diff(+got -want) %v", tt.input, diff)
		}
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		universe Universe
		g1, g2   Galaxy
		want     int
	}{
		{

			Universe{
				DistanceX: []int{1, 1, 1},
				DistanceY: []int{1, 1, 1},
			},
			Galaxy{0, 0},
			Galaxy{2, 0},
			2,
		},
		{

			Universe{
				DistanceX: []int{1, 2, 1},
				DistanceY: []int{1, 1, 1},
			},
			Galaxy{0, 0},
			Galaxy{2, 0},
			3,
		},
		{

			Universe{
				DistanceX: []int{1, 1, 1},
				DistanceY: []int{1, 1, 1},
			},
			Galaxy{0, 0},
			Galaxy{2, 2},
			4,
		},
		{

			Universe{
				DistanceX: []int{1, 2, 1},
				DistanceY: []int{1, 2, 1},
			},
			Galaxy{0, 0},
			Galaxy{2, 2},
			6,
		},
	}
	for _, tt := range tests {
		got := tt.universe.Distance(&tt.g1, &tt.g2)
		if got != tt.want {
			t.Errorf("universe %+v distance %v %v: got %d, want %d", tt.universe, tt.g1, tt.g2, got, tt.want)
		}
	}
}

var defaultInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			defaultInput,
			374,
		},
	}

	for _, tt := range tests {
		got, err := part1(tt.input)
		if err != nil {
			t.Fatalf("part1(%q) error %v, want %v", tt.input, err, nil)
		}

		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Errorf("part1(%q) diff(+got -want) %v", tt.input, diff)
		}
	}
}

func TestSumWithScale(t *testing.T) {
	tests := []struct {
		input string
		scale int
		want  int
	}{
		{
			defaultInput,
			10,
			1030,
		},
		{
			defaultInput,
			100,
			8410,
		},
	}

	for _, tt := range tests {
		got, err := answerWithScale(tt.input, tt.scale)
		if err != nil {
			t.Fatalf("answerWithScale(%q) error %v, want %v", tt.input, err, nil)
		}

		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Errorf("answerWithScale(%q) diff(+got -want) %v", tt.input, diff)
		}
	}
}
