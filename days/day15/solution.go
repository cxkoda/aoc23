package day15

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/cxkoda/aoc23/days"
)

func hash(s string, current int) int {
	if s == "" {
		return current
	}

	current += int(s[0])
	current = (current * 17) % 256

	return hash(s[1:], current)
}

func Hash(s string) int {
	return hash(s, 0)
}

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	Lenses []Lens
}

func (b *Box) AddLens(lens Lens) {
	for i, l := range b.Lenses {
		if l.Label == lens.Label {
			b.Lenses[i] = lens
			return
		}
	}

	b.Lenses = append(b.Lenses, lens)
}

func (b *Box) RemoveLens(label string) {
	for i, l := range b.Lenses {
		if l.Label == label {
			b.Lenses = append(b.Lenses[:i], b.Lenses[i+1:]...)
			break
		}
	}
}

type Boxes [256]Box

func (b *Boxes) Execute(action string) error {
	len := len(action)

	if action[len-1] == '-' {
		// deletion
		box := &b[Hash(action[:len-1])]
		box.RemoveLens(action[:len-1])
		return nil
	}

	box := &b[Hash(action[:len-2])]
	focalLength, err := strconv.Atoi(action[len-1:])
	if err != nil {
		return fmt.Errorf("invalid focal length %v", action[len-1:])
	}

	box.AddLens(Lens{
		Label:       action[:len-2],
		FocalLength: focalLength,
	})
	return nil
}

func (b *Boxes) Sum() int {
	var sum int
	for i, box := range b {
		for j, lens := range box.Lenses {
			sum += lens.FocalLength * (i + 1) * (j + 1)
		}
	}
	return sum
}

//go:embed input
var input Challenge

type Challenge string

func init() {
	days.MustRegister(15, input)
}

func (in Challenge) Part1() (int, error) {
	var sum int
	for _, s := range strings.Split(string(in), ",") {
		sum += Hash(s)
	}

	return sum, nil
}

func (in Challenge) Part2() (int, error) {
	bs := Boxes{}
	for _, s := range strings.Split(string(in), ",") {
		if err := bs.Execute(s); err != nil {
			return 0, err
		}
	}

	return bs.Sum(), nil
}
