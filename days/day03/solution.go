package day03

import (
	_ "embed"
	"math"
	"strings"

	"github.com/cxkoda/aoc23/days"
	"github.com/cxkoda/aoc23/days/day01"
)

type number struct {
	val    int
	active bool
}

func newGrid(s string) grid {
	var g grid
	for _, r := range strings.Split(s, "\n") {
		g = append(g, []byte(r))
	}

	return g
}

type grid [][]byte

type Direction int

const (
	Left  Direction = -1
	Right           = +1
)

func getNumber(s []byte, i int, clear bool) (int, bool) {
	v, ok := day01.ParseDigit(s[i])
	if !ok {
		return 0, false
	}

	if l, ok := getAdjacentNumber(s, i, Left, clear); ok {
		v += 10 * l
	}

	if r, ok := getAdjacentNumber(s, i, Right, clear); ok {
		v = int(math.Pow10(int(math.Log10(float64(r)))+1))*v + r
	}

	return v, true
}

func getAdjacentNumber(s []byte, i int, dir Direction, clear bool) (int, bool) {
	i += int(dir)
	if i < 0 || i >= len(s) {
		return 0, false
	}

	v, ok := day01.ParseDigit(s[i])
	if !ok {
		return 0, false
	}

	if clear {
		s[i] = '.'
	}

	w, ok := getAdjacentNumber(s, i, dir, clear)
	if !ok {
		return v, true
	}

	if dir == Left {
		return 10*w + v, true
	}

	return int(math.Pow10(int(math.Log10(float64(w)))+1))*v + w, true
}

func isSymbol(c byte) bool {
	if c == '.' {
		return false
	}

	if _, ok := day01.ParseDigit(c); ok {
		return false
	}

	return true
}

func (g grid) searchNumbersWithSymbols(clear bool) []int {
	var nums []int
	appendIfNum := func(x, y int) {
		if num, ok := getNumber(g[y], x, clear); ok {
			nums = append(nums, num)
		}
	}

	for y, row := range g {
		for x, c := range row {
			if !isSymbol(c) {
				continue
			}

			if y > 0 {
				appendIfNum(x, y-1)
			}

			if y < len(g)-1 {
				appendIfNum(x, y+1)
			}

			if x > 0 {
				appendIfNum(x-1, y)
			}

			if x < len(row)-1 {
				appendIfNum(x+1, y)
			}

			if y > 0 && x > 0 {
				appendIfNum(x-1, y-1)
			}

			if y > 0 && x < len(row)-1 {
				appendIfNum(x+1, y-1)
			}

			if y < len(g)-1 && x > 0 {
				appendIfNum(x-1, y+1)
			}

			if y < len(g)-1 && x < len(row)-1 {
				appendIfNum(x+1, y+1)
			}

		}
	}
	return nums
}

//go:embed input
var input Challenge

type Challenge string

func init() {
	days.MustRegister(3, input)
}

func (in Challenge) Part1() (int, error) {
	nums := newGrid(string(in)).searchNumbersWithSymbols(true)
	var sum int
	for _, x := range nums {
		sum += x
	}

	return sum, nil
}

func (in Challenge) Part2() (int, error) {
	return 0, nil
}
