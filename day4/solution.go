package day4

import (
	"iter"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) Move(dir Direction) Point {
	return Point{p.X + dir.Dx, p.Y + dir.Dy}
}

type Direction struct {
	Dx int
	Dy int
}

var Directions = []Direction{
	{1, 1}, {1, 0}, {1, -1},
	{0, 1}, {0, -1},
	{-1, 1}, {-1, 0}, {-1, -1},
}

type Cell int

const (
	OutOfBound = iota
	Free
	Roll
)

type Grid [][]Cell

func (gr Grid) IsInside(point Point) bool {
	x, y := point.X, point.Y
	return !(x < 0 || x >= len(gr) || y < 0 || y >= len(gr[0]))
}

func (gr Grid) Get(point Point) Cell {
	if !gr.IsInside(point) {
		return OutOfBound
	}
	return gr[point.X][point.Y]
}

func (gr Grid) Neighbors(point Point) iter.Seq[Cell] {
	if !gr.IsInside(point) {
		panic("Point is outside")
	}

	return func(yield func(Cell) bool) {
		for _, dir := range Directions {
			point := gr.Get(point.Move(dir))
			if !yield(point) {
				break
			}
		}
	}
}

func (gr Grid) CellsSeq() iter.Seq2[Point, Cell] {
	return func(yield func(Point, Cell) bool) {
		for x, row := range gr {
			for y, cell := range row {
				if !yield(Point{x, y}, cell) {
					return
				}
			}
		}
	}
}

func (gr Grid) RemoveRoll(point Point) {
	if gr.Get(point) != Roll {
		panic("point is not a roll")
	}
	gr[point.X][point.Y] = Free
}

func parseCell(val rune) Cell {
	switch val {
	case '.':
		return Free
	case '@':
		return Roll
	default:
		panic("Unknown cell")
	}
}

func Parse(filename string) Grid {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	grid := Grid{}
	for line := range strings.SplitSeq(string(data), "\n") {
		row := []Cell{}
		for _, val := range line {
			row = append(row, parseCell(val))
		}
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}

	return grid
}

func GetRollsAround(grid Grid, point Point) int {
	result := 0
	for cell := range grid.Neighbors(point) {
		if cell == Roll {
			result += 1
		}
	}
	return result
}

func Part1(grid Grid) int {
	return len(GetRollsToRemove(grid))
}

func GetRollsToRemove(grid Grid) []Point {
	result := map[Point]struct{}{}

	for point, cell := range grid.CellsSeq() {
		if cell == Roll && GetRollsAround(grid, point) < 4 {
			result[point] = struct{}{}
		}
	}

	rv := []Point{}

	for point := range result {
		rv = append(rv, point)
	}
	return rv
}

func Part2(input Grid) int {
	result := 0

	for q := GetRollsToRemove(input); len(q) > 0; q = GetRollsToRemove(input) {
		println(result, q)
		result += len(q)
		for _, p := range q {
			if input.IsInside(p) {
				input.RemoveRoll(p)
			}
		}
	}
	return result
}
