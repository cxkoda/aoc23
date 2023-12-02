package main

import "testing"

func TestDecodeCalibration(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, tt := range tests {
		got, err := decodeCalibration(tt.input)
		if err != nil {
			t.Errorf("decodeCalibration(%q) error %v, want %v", tt.input, err, nil)
		}

		if got != tt.want {
			t.Errorf("decodeCalibration(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestTotalCalibration(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142},
	}
	for _, tt := range tests {
		got, err := totalCalibration(tt.input)
		if err != nil {
			t.Errorf("totalCalibration(%q) error %v, want %v", tt.input, err, nil)
		}

		if got != tt.want {
			t.Errorf("totalCalibration(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
