package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input []int

func Parse(filename string) Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	var result Input
	for _, v := range strings.Split(string(data), "\n") {
		if len(v) == 0 {
			break
		}
		dir := 1
		if v[0] == 'L' {
			dir = -1
		}
		val, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err.Error())
		}
		result = append(result, dir*val)
	}
	return result
}

func Part1(input Input) int {
	cur := 50
	result := 0

	for _, step := range input {
		cur = (cur + 100 + step) % 100
		if cur == 0 {
			result += 1
		}
	}
	return result
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Part2(input Input) int {
	cur := 50
	result := 0

	for _, step := range input {
		for i := 0; i < abs(step); i++ {
			if step < 0 {
				cur -= 1
			} else {
				cur += 1
			}
			cur = cur % 100
			if cur == 0 {
				result += 1
			}
		}
		fmt.Println(cur)
	}
	return result
}
