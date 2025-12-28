package day7

import (
	"fmt"
	"os"
	"strings"
)

type Input []string

type Beams struct {
	pos []int
	length int
	sum int
}

func (bs *Beams) Add(pos int, num int)  {
	if bs.pos[pos] == 0 {
		bs.length += 1
	}

	bs.sum += num
	bs.pos[pos] += num
}

func (bs *Beams) Remove(pos int)  {
	if bs.pos[pos] > 0 {
		bs.length -= 1
	}
	bs.sum -= bs.pos[pos]
	bs.pos[pos] = 0
}

func (bs *Beams) Split(pos int) bool {
	if bs.pos[pos] == 0 {
		return false
	}
	bs.Add(pos-1, bs.pos[pos])
	bs.Add(pos+1, bs.pos[pos])
	bs.Remove(pos)
	return true
}

func (bs *Beams) Len() int {
	return bs.length
}

func (bs *Beams) String() string {
	return fmt.Sprint(bs.pos)
}

func (bs *Beams) Sum() int {
	return bs.sum
}

func NewBeams(N int) *Beams {
	return &Beams{
		make([]int, N),
		0,
		0,
	}
}

func Parse(filename string) Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func Part1(input Input) int {
	beams := NewBeams(len(input))
	result := 0
	for _, line := range input {
		for i, r := range line {
			switch r {
			case 'S':
				beams.Add(i, 1)
			case '^':
				if beams.Split(i) {
					result += 1
				}
			}
		}
	}
	return result
}

func Part2(input Input) int {
	beams := NewBeams(len(input))
	for _, line := range input {
		for i, r := range line {
			switch r {
			case 'S':
				beams.Add(i, 1)
			case '^':
				beams.Split(i)
			}
		}
	}
	return beams.Sum()
}

