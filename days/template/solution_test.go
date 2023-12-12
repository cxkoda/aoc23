package template

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		// {
		// 	``,
		// 	0,
		// },
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
		// {
		// 	``,
		// 	0,
		// },
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
