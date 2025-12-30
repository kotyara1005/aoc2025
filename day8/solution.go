package day8

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/kotyara1005/aoc2025/utils"
)

type Point3D struct {
	X int
	Y int
	Z int
}

func NewPointFromString(line string) Point3D {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		panic("bad point")
	}
	return Point3D{
		utils.Atoi(parts[0]),
		utils.Atoi(parts[1]),
		utils.Atoi(parts[2]),
	}
}

func GetDistance(a, b Point3D) float64 {
	return math.Sqrt(
		math.Pow(float64(a.X-b.X), 2) +
			math.Pow(float64(a.Y-b.Y), 2) +
			math.Pow(float64(a.Z-b.Z), 2),
	)
}

func Parse(filename string) []Point3D {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	result := []Point3D{}

	for line := range strings.Lines(string(data)) {
		result = append(result, NewPointFromString(strings.Trim(line, "\n")))
	}
	return result
}

type dist struct {
	Dst float64
	P1  int
	P2  int
}

func getdistances(points []Point3D) []dist {
	result := []dist{}

	for i, p1 := range points {
		for j, p2 := range points {
			if j <= i {
				continue
			}
			d := GetDistance(p1, p2)
			result = append(result, dist{Dst: d, P1: i, P2: j})
		}
	}
	return result
}

func Part1(points []Point3D, iters int) int {
	uf := utils.NewUnionFind[int]()
	for i := range points {
		uf.Add(i)
	}

	distances := getdistances(points)
	fmt.Println(len(points))
	slices.SortFunc(distances, func(a, b dist) int {
		if a.Dst < b.Dst {
			return -1
		}
		if a.Dst > b.Dst {
			return 1
		}
		return 0
	})
	distances = distances[:iters]
	fmt.Println(len(distances))

	for _, d := range distances {
		fmt.Println("Joined", d, uf.Join(d.P1, d.P2))
	}

	processed := map[int]struct{}{}

	sizes := []int{}
	for i := range points {
		key := uf.Find(i)
		if _, prs := processed[key]; prs {
			continue
		}
		processed[key] = struct{}{}
		sizes = append(sizes, uf.Size(key))
	}
	sort.Slice(sizes, func(i, j int) bool {return sizes[i] > sizes[j]})
	fmt.Println(sizes)
	return sizes[0] * sizes[1] * sizes[2]
}

func Part2(points []Point3D) int {
	uf := utils.NewUnionFind[int]()
	for i := range points {
		uf.Add(i)
	}

	distances := getdistances(points)
	sort.Slice(distances, func(i, j int) bool {return distances[i].Dst < distances[j].Dst})

	for _, d := range distances {
		uf.Join(d.P1, d.P2)
		
		if uf.Len() == 1 {
			p1 := points[d.P1]
			p2 := points[d.P2]
			fmt.Println(p1, p2)
			return p1.X * p2.X
		}
	}

	return 0
}
