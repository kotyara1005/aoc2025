package day5

import (
	"os"
	"sort"
	"strings"

	"github.com/kotyara1005/aoc2025/utils"
)

type Input struct {
	utils.Intervals
	Ids []int
}

func parseIntervals(data string) utils.Intervals {
	result := utils.Intervals{}
	for line := range strings.Lines(data) {
		line = strings.TrimRight(line, "\n")
		elems := strings.Split(line, "-")
		result = append(result, utils.Interval{utils.Atoi(elems[0]), utils.Atoi(elems[1])})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	rv := utils.Intervals{result[0]}
	for _, it := range result[1:] {
		if it[0] <= rv[len(rv)-1][1] + 1 {
			rv[len(rv)-1][1] = max(rv[len(rv)-1][1], it[1])
		} else {
			rv = append(rv, it)
		}
	}
	return rv
}

func parseIds(data string) []int {
	result := []int{}
	for line := range strings.Lines(data) {
		result = append(result, utils.Atoi(strings.Trim(line, "\n")))
	}
	return result
}

func Parse(filename string) Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	parts := strings.Split(string(data), "\n\n")
	return Input{
		Intervals: parseIntervals(parts[0]),
		Ids:       parseIds(parts[1]),
	}
}

func Part1(input Input) int {
	result := 0
	for _, id := range input.Ids {
		// println(id, input.Intervals.Contains(id))
		if input.Intervals.Contains(id) {
			result += 1
		}
	}
	return result
}

func Part2(input Input) int {
	result := 0
	for _, it := range input.Intervals {
		result += it[1] - it[0] + 1
	}
	return result
}
