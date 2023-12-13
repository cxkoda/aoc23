package template

import (
	_ "embed"
	"strings"

	"github.com/cxkoda/aoc23/days"
)

func NewMap(s string) Map {
	var m Map
	for _, line := range strings.Split(s, "\n") {
		var row []bool
		for _, c := range line {
			row = append(row, c == '#')
		}
		m = append(m, row)
	}
	return m
}

type Map [][]bool

func (m Map) NumRows() int {
	return len(m)
}

func (m Map) NumCols() int {
	return len(m[0])
}

type Direction int

const (
	Vertical Direction = iota
	Horizontal
)

type Reflection struct {
	Direction Direction
	Position  int
}

func (r Reflection) reflect(x, y int) (int, int) {
	switch r.Direction {
	case Horizontal:
		return x, 2*r.Position - y - 1
	case Vertical:
		return 2*r.Position - x - 1, y
	}

	panic("unreachable")
}

func (r Reflection) Equal(other Reflection) bool {
	return r.Direction == other.Direction && r.Position == other.Position
}

func (m Map) CheckReflection(r Reflection) bool {
	for x := 0; x < m.NumCols(); x++ {
		for y := 0; y < m.NumRows(); y++ {
			xr, yr := r.reflect(x, y)
			if xr < 0 || xr >= m.NumCols() || yr < 0 || yr >= m.NumRows() {
				continue
			}

			if m[y][x] != m[yr][xr] {
				return false
			}
		}
	}
	return true
}

func (m Map) PossibleReflections() []Reflection {
	var rs []Reflection

	for p := 1; p < m.NumRows(); p++ {
		rs = append(rs, Reflection{Direction: Horizontal, Position: p})
	}

	for p := 1; p < m.NumCols(); p++ {
		rs = append(rs, Reflection{Direction: Vertical, Position: p})
	}

	return rs
}

func (m Map) FindReflections() []Reflection {
	var reflections []Reflection

	rs := m.PossibleReflections()
	for _, r := range rs {
		if m.CheckReflection(r) {
			reflections = append(reflections, r)
		}
	}

	return reflections
}

//go:embed input
var input Challenge

type Challenge string

func init() {
	days.MustRegister(13, input)
}

func (in Challenge) Part1() (int, error) {
	var ms []Map
	for _, s := range strings.Split(string(in), "\n\n") {
		ms = append(ms, NewMap(s))
	}

	var sum int
	for _, m := range ms {
		rs := m.FindReflections()
		for _, r := range rs {
			v := r.Position
			if r.Direction == Horizontal {
				v *= 100
			}
			sum += v
		}
	}

	return sum, nil
}

func filter(input, exclude []Reflection) []Reflection {
	ex := make(map[Reflection]bool)
	for _, e := range exclude {
		ex[e] = true
	}

	var rs []Reflection
	for _, r := range input {
		if !ex[r] {
			rs = append(rs, r)
		}
	}
	return rs
}

func (in Challenge) Part2() (int, error) {
	var ms []Map
	for _, s := range strings.Split(string(in), "\n\n") {
		ms = append(ms, NewMap(s))
	}

	var sum int
	for _, m := range ms {
		rs := m.FindReflections()

	MutationLoop:
		for y, row := range m {
			for x, c := range row {
				m[y][x] = !c

				rsNew := m.FindReflections()
				rsNew = filter(rsNew, rs)
				if len(rsNew) > 0 {
					rs = rsNew
					break MutationLoop
				}

				m[y][x] = c
			}
		}

		for _, r := range rs {
			v := r.Position
			if r.Direction == Horizontal {
				v *= 100
			}
			sum += v
		}
	}

	return sum, nil
}
