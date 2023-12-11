package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	_ "embed"
)

type Galaxy struct {
	X, Y int
}

type Universe struct {
	Galaxies             []*Galaxy
	DistanceX, DistanceY []int
}

func parseUniverse(s string, emptyScale int) (*Universe, error) {
	rows := strings.Split(s, "\n")

	u := Universe{
		DistanceX: make([]int, len(rows[0])),
		DistanceY: make([]int, len(rows)),
	}

	for y := range u.DistanceY {
		u.DistanceY[y] = emptyScale
	}
	for x := range u.DistanceX {
		u.DistanceX[x] = emptyScale
	}

	for y, row := range rows {
		if len(row) != len(rows[0]) {
			return nil, fmt.Errorf("invalid row length %d, want %d", len(row), len(rows[0]))
		}
		for x, v := range row {
			if v == '#' {
				u.Galaxies = append(u.Galaxies, &Galaxy{x, y})
				u.DistanceX[x] = 1
				u.DistanceY[y] = 1
			}
		}
	}

	return &u, nil
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (u *Universe) Distance(g1, g2 *Galaxy) int {
	var d int

	xL, xR := minMax(g1.X, g2.X)
	for x := xL; x < xR; x++ {
		d += u.DistanceX[x]
	}

	yL, yR := minMax(g1.Y, g2.Y)
	for y := yL; y < yR; y++ {
		d += u.DistanceY[y]
	}

	return d
}

type Pair struct {
	G1, G2 *Galaxy
}

func getPairs(gs []*Galaxy) []*Pair {
	var pairs []*Pair
	for i, g1 := range gs {
		for _, g2 := range gs[i+1:] {
			pairs = append(pairs, &Pair{g1, g2})
		}
	}
	return pairs
}

func (u *Universe) SumDistances() int {
	var d int
	for _, p := range getPairs(u.Galaxies) {
		d += u.Distance(p.G1, p.G2)
	}
	return d
}

//go:embed input
var input string

func answerWithScale(input string, scale int) (int, error) {
	u, err := parseUniverse(input, scale)
	if err != nil {
		return 0, err
	}

	return u.SumDistances(), nil
}

func part1(input string) (int, error) {
	return answerWithScale(input, 2)
}

func part2(input string) (int, error) {
	return answerWithScale(input, 1000000)
}

func run() error {
	ans1, err := part1(input)
	if err != nil {
		return err
	}
	fmt.Println("Part 1:", ans1)

	ans2, err := part2(input)
	if err != nil {
		return err
	}
	fmt.Println("Part 2:", ans2)

	return nil
}

func main() {
	now := time.Now()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	diff := time.Since(now)
	fmt.Println("Time:", diff)
}
