package day03

import (
	"testing"
)

func TestGetNumber(t *testing.T) {
	tests := []struct {
		s      string
		i      int
		want   int
		wantOk bool
	}{
		{
			s:      ".137.",
			i:      0,
			wantOk: false,
		},
		{
			s:      ".137.",
			i:      1,
			want:   137,
			wantOk: true,
		},
		{
			s:      ".137.",
			i:      2,
			want:   137,
			wantOk: true,
		},
		{
			s:      ".137.",
			i:      3,
			want:   137,
			wantOk: true,
		},
		{
			s:      ".137.",
			i:      4,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		got, ok := getNumber([]byte(tt.s), tt.i, false)
		if ok != tt.wantOk {
			t.Errorf("getNumber(%s, %d, ..), ok=%t, want %t", tt.s, tt.i, ok, tt.wantOk)
		}

		if got != tt.want {
			t.Errorf("getNumber(%s, %d, ..)=%d, want %d", tt.s, tt.i, got, tt.want)
		}
	}

}

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			4361,
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
