package day9

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/kotyara1005/aoc2025/utils"
	"github.com/kotyara1005/aoc2025/utils/itertools"
)

type Point struct {
	X int
	Y int
}

func NewPoint(line string) Point {
	parts := strings.Split(strings.Trim(line, "\n"), ",")
	if len(parts) != 2 {
		panic("bad point")
	}
	return Point{utils.Atoi(parts[0]), utils.Atoi(parts[1])}
}

type Rectangle struct {
	P1 Point
	P2 Point
}

func (rec *Rectangle) Square() int {
	return (max(rec.P1.X, rec.P2.X) - min(rec.P1.X, rec.P2.X) + 1) * (max(rec.P1.Y, rec.P2.Y) - min(rec.P1.Y, rec.P2.Y) + 1)
}

func (rec *Rectangle) Perimeter() []Point {
	var xMin, xMax, yMin, yMax int
	xMin = min(rec.P1.X, rec.P2.X)
	xMax = max(rec.P1.X, rec.P2.X)
	yMin = min(rec.P1.Y, rec.P2.Y)
	yMax = max(rec.P1.Y, rec.P2.Y)

	rv := []Point{}
	for x := xMin; x <= xMax; x++ {
		rv = append(rv, Point{x, yMin}, Point{x, yMax})
	}
	for y := yMin; y <= yMax; y++ {
		rv = append(rv, Point{xMin, y}, Point{xMax, y})
	}
	return rv
}

func NewRectangle(p1, p2 Point) Rectangle {
	return Rectangle{p1, p2}
}

func Parse(filename string) []Point {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	result := []Point{}
	for line := range strings.Lines(string(data)) {
		result = append(result, NewPoint(line))
	}
	return result
}

func Part1(points []Point) int {
	result := NewRectangle(Point{0, 0}, Point{0, 0})
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			loc := NewRectangle(p1, p2)
			if loc.Square() > result.Square() {
				result = loc
			}
		}
	}
	return result.Square()
}

type Boundaries struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

func (bn Boundaries) Contains(point Point) bool {
	return bn.Left <= point.X && point.X <= bn.Right && bn.Bottom <= point.Y && point.Y <= bn.Top
}

func NewBoundariesFromPoints(points []Point) Boundaries {
	rv := Boundaries{}
	rv.Left = slices.MinFunc(points, func(p1, p2 Point) int {
		return cmp.Compare(int64(p1.X), int64(p2.X))
	}).X - 1
	rv.Right = slices.MaxFunc(points, func(p1, p2 Point) int {
		return cmp.Compare(int64(p1.X), int64(p2.X))
	}).X + 1
	rv.Bottom = slices.MinFunc(points, func(p1, p2 Point) int {
		return cmp.Compare(int64(p1.Y), int64(p2.Y))
	}).Y - 1
	rv.Top = slices.MaxFunc(points, func(p1, p2 Point) int {
		return cmp.Compare(int64(p1.Y), int64(p2.Y))
	}).Y + 1
	return rv
}

type Edge [2]Point

func (ed Edge) Contains(point Point) bool {
	isHorizontal := ed[0].X == ed[1].X

	if isHorizontal {
		if point.X != ed[0].X {
			return false
		}
		return min(ed[0].Y, ed[1].Y) <= point.Y && point.Y <= max(ed[0].Y, ed[1].Y)
	} else {
		if point.Y != ed[0].Y {
			return false
		}
		return min(ed[0].X, ed[1].X) <= point.X && point.X <= max(ed[0].X, ed[1].X)
	}
}

func NewEdge(p1, p2 Point) Edge {
	if p1.X != p2.X && p1.Y != p2.Y {
		panic("bad edge")
	}
	return Edge{p1, p2}
}

type Polygon []Edge

func (eds Polygon) Contains(point Point) bool {
	for _, ed := range eds {
		if ed.Contains(point) {
			return true
		}
	}
	return false
}

func NewPolygonFromChain(points []Point) Polygon {
	pl := Polygon{}
	for p1, p2 := range itertools.WithPrev(slices.Values(points)) {
		pl = append(pl, NewEdge(p1, p2))
	}
	pl = append(pl, NewEdge(points[0], points[len(points)-1]))
	return pl
}

func GetNeghbours(point Point) []Point {
	return []Point{
		{point.X + 1, point.Y},
		{point.X - 1, point.Y},
		{point.X, point.Y + 1},
		{point.X, point.Y - 1},
	}
}

func FloodFill(boundaries Boundaries, edges Polygon) map[Point]struct{} {
	empty := sync.Map{}
	isValid := func(p Point) bool {
		return boundaries.Contains(p) && !edges.Contains(p)
	}
	visited := func(p Point) bool {
		_, prs := empty.Load(p)
		return prs
	}

	startingPoints := []Point{
		{boundaries.Left, boundaries.Bottom},
		{boundaries.Left, boundaries.Top},
		{boundaries.Right, boundaries.Bottom},
		{boundaries.Right, boundaries.Top},
	}

	var processed int64 = 0
	wg := sync.WaitGroup{}

	bfs := func(start Point) {
		q := []Point{start}
		empty.Store(start, struct{}{})

		for len(q) > 0 {
			fmt.Println(len(q), atomic.LoadInt64(&processed))
			nextQ := []Point{}
			for _, point := range q {
				for _, neighbour := range GetNeghbours(point) {
					if !isValid(neighbour) || visited(neighbour) {
						continue
					}
					empty.Store(neighbour, struct{}{})
					atomic.AddInt64(&processed, 1)
					nextQ = append(nextQ, neighbour)
				}
			}
			q = nextQ
		}
		wg.Done()
	}

	for _, point := range startingPoints {
		wg.Add(1)
		go bfs(point)
	}
	wg.Wait()

	result := map[Point]struct{}{}

	empty.Range(func(key, value any) bool {
		result[key.(Point)] = struct{}{}
		return false
	})
	return result
}

func CheckRectangleInPolygon(rec Rectangle, outside map[Point]struct{}) bool {
	for _, point := range rec.Perimeter() {
		_, prs := outside[point]
		if prs {
			return false
		}
	}
	return true
}

// func PrintGrid(boundaries Boundaries, outside map[Point]struct{}) {
// 	builder := strings.Builder{}

// 	for y := yMin; y <= yMax; y++ {

// 	}

// }

func Part2(points []Point) int {
	// find boundaries
	boundaries := NewBoundariesFromPoints(points)
	fmt.Println(boundaries)
	polygon := NewPolygonFromChain(points)
	// fill flood = bfs and bounce of the edges
	outside := FloodFill(boundaries, polygon)

	fmt.Println(len(outside))
	// check that edges of rectangle are in polygon
	result := NewRectangle(Point{0, 0}, Point{0, 0})
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			loc := NewRectangle(p1, p2)
			if !CheckRectangleInPolygon(loc, outside) {
				continue
			}
			if loc.Square() > result.Square() {
				result = loc
			}
		}
	}
	return result.Square()
}
