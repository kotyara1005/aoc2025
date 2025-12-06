package day3

import (
	"math"
	"os"
	"slices"
	"strings"

	"github.com/kotyara1005/aoc2025/utils"
)

type Batary []int

func (bt Batary) MaxJoltage() int {
	result := 0
	stack := []int{}

	for _, val := range bt {
		for len(stack) > 0 && stack[len(stack)-1] < val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}
	slices.Reverse(stack)

	left := 0
	for _, val := range bt[:len(bt)-1] {
		if len(stack) > 0 && stack[len(stack)-1] == val {
			stack = stack[:len(stack)-1]
		}
		left = max(left, val)
		right := stack[len(stack)-1]
		result = max(result, left*10+right)
	}
	return result
}


func (bt Batary) buildGraph() [][10]int {
	result := [][10]int{}
	for idx := range bt {
		cur := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
		leftToFill := 9
		for j, val := range bt[idx+1:] {
			if cur[val] == -1 {
				cur[val] = idx + 1 + j
				leftToFill -= 1
			}
			if leftToFill == 0 {
				break
			}
		}
		result = append(result, cur)
	}
	return result
}

func (bt Batary) dfs(digit int, pos int, init [10]int, graph [][10]int) int {
	if digit == 0 {
		return 0
	}

	node := init
	if pos != -1 {
		node = graph[pos]
	}

	for i := 9; i > 0; i-- {
		nxt := node[i]
		if nxt == -1 || len(bt)-nxt < digit {
			continue
		}
		return i*int(math.Pow10(digit-1)) + bt.dfs(digit-1, nxt, init, graph)
	}

	return 0
}

func (bt Batary) MaxJoltage12() int {
	graph := bt.buildGraph()
	init := graph[0]
	init[bt[0]] = 0
	return bt.dfs(12, -1, init, graph)
}

type Bataries []Batary

func Parse(filename string) Bataries {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	result := Bataries{}
	for line := range strings.Lines(string(data)) {
		bt := Batary{}
		for _, char := range line {
			if char == '\n' {
				continue
			}
			bt = append(bt, utils.Atoi(string([]rune{char})))
		}
		result = append(result, bt)
	}
	return result
}

func Part1(bataries Bataries) int {
	result := 0
	for _, batary := range bataries {
		result += batary.MaxJoltage()
	}
	return result
}

func Part2(bataries Bataries) int {
	result := 0
	for _, batary := range bataries {
		// fmt.Println(idx)
		result += batary.MaxJoltage12()
	}
	return result
}
