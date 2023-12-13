package template

import (
	"reflect"
	"testing"
)

func TestCheckReflection(t *testing.T) {
	tests := []struct {
		m    Map
		r    Reflection
		want bool
	}{
		{
			Map{
				{true, false},
				{true, false},
			},
			Reflection{
				Direction: Horizontal,
				Position:  1,
			},
			true,
		},
		{
			Map{
				{true, false},
				{true, false},
			},
			Reflection{
				Direction: Vertical,
				Position:  1,
			},
			false,
		},
	}

	for _, tt := range tests {
		if got := tt.m.CheckReflection(tt.r); got != tt.want {
			t.Errorf("%v.CheckReflection(%v) = %v, want %v", tt.m, tt.r, got, tt.want)
		}
	}
}

func TestFindReflections(t *testing.T) {
	tests := []struct {
		m    Map
		want []Reflection
	}{
		{
			Map{
				{true, false},
				{true, false},
			},
			[]Reflection{
				{
					Direction: Horizontal,
					Position:  1,
				},
			},
		},
		{
			Map{
				{false, true, true},
				{false, true, true},
			},
			[]Reflection{
				{
					Direction: Horizontal,
					Position:  1,
				},
				{
					Direction: Vertical,
					Position:  2,
				},
			},
		},
		{
			NewMap(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`),
			[]Reflection{
				{
					Direction: Vertical,
					Position:  5,
				},
			},
		},
		{
			// after modification
			NewMap(`..##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`),
			[]Reflection{
				{
					Direction: Horizontal,
					Position:  3,
				},
				{
					Direction: Vertical,
					Position:  5,
				},
			},
		},
		{
			NewMap(`#...##..#
#...##..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`),
			[]Reflection{
				{
					Direction: Horizontal,
					Position:  1,
				},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.m.FindReflections(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v.FindReflections() = %v, want %v", tt.m, got, tt.want)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
			405,
		},
	}
	for _, tt := range tests {
		challenge := Challenge(tt.input)
		got, err := challenge.Part1()
		if err != nil {
			t.Fatalf("%T.Part1() error %v, want %v", challenge, err, nil)
		}

		if got != tt.want {
			t.Errorf("%T.Part1() = %v, want %v", challenge, got, tt.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
			400,
		},
	}
	for _, tt := range tests {
		challenge := Challenge(tt.input)
		got, err := challenge.Part2()
		if err != nil {
			t.Fatalf("%T.Part2() error %v, want %v", challenge, err, nil)
		}

		if got != tt.want {
			t.Errorf("%T.Part2() = %v, want %v", challenge, got, tt.want)
		}
	}
}
