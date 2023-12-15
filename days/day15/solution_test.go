package day15

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			"",
			0,
		},
		{
			"rn=1",
			30,
		},
		{
			"cm-",
			253,
		},
		{
			"qp=3",
			97,
		},
	}

	for _, tt := range tests {
		got := Hash(tt.input)
		if got != tt.want {
			t.Errorf("Hash(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
			1320,
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
			`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
			145,
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
