package days

import (
	"fmt"
	"io"
	"sort"
	"time"
)

type Solution interface {
	Part1() (int, error)
	Part2() (int, error)
}

func register(day int, sol Solution) error {
	if _, ok := solutions[day]; ok {
		return fmt.Errorf("solution for day %d already exists", day)
	}
	solutions[day] = sol
	return nil
}

func MustRegister(day int, sol Solution) {
	if err := register(day, sol); err != nil {
		// replace with fatalf
		panic(err)
	}
}

func EvaluateAll(w io.Writer) error {
	ds := make([]int, 0, len(solutions))
	for d := range solutions {
		ds = append(ds, d)
	}

	sort.IntSlice(ds).Sort()
	for _, d := range ds {
		sol := solutions[d]
		for p, f := range []func() (int, error){sol.Part1, sol.Part2} {
			now := time.Now()
			ans, err := f()
			if err != nil {

				return fmt.Errorf("%T() error %v", f, err)
			}
			elapsed := time.Since(now)
			if _, err := fmt.Fprintf(w, "Day %d, part %d [%v]: %d\n", d, p, elapsed, ans); err != nil {
				return fmt.Errorf("fmt.Fprintf(%T, ...) error %v", w, err)
			}
		}
	}

	return nil
}

var solutions map[int]Solution

func init() {
	solutions = make(map[int]Solution)
}
