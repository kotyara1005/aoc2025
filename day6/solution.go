package day6

import (
	"os"
	"strconv"
	"strings"
)

//go:generate stringer -type=Operation
type Operation rune

const (
	OpSum Operation = '+'
	OpMul Operation = '*'
)

func OperationFromRune(rn rune) Operation {
	switch Operation(rn) {
	case OpSum:
		return OpSum
	case OpMul:
		return OpMul
	default:
		panic("unknown operation")
	}
}

type Input struct {
	Numbers    [][]int
	Operations []Operation
}

func Parse(filename string) Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(data), "\n")
	// fmt.Println(len(lines))
	result := Input{}

	for _, line := range lines[:len(lines)-2] {
		// fmt.Println(i)
		nums := []int{}
		for val := range strings.SplitSeq(line, " ") {
			num, err := strconv.Atoi(val)
			if err == nil {
				nums = append(nums, num)
			}
		}
		result.Numbers = append(result.Numbers, nums)
	}

	for _, val := range lines[len(lines)-2] {
		if val == ' ' {
			continue
		}
		result.Operations = append(result.Operations, OperationFromRune(val))
	}

	return result
}

func ParsePart2(filename string) Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(data), "\n")

	result := Input{}
	numbers := [][]rune{}
	for i := 0; i < len(lines)-2; i++ {
		numbers = append(numbers, []rune(lines[i]))
	}

	result.Numbers = append(result.Numbers, []int{})
	for i := 0; i < len(numbers[0]); i++ {
		line := []rune{}
		for j := 0; j < len(numbers); j++ {
			line = append(line, numbers[j][i])
		}
		num, err := strconv.Atoi(strings.TrimSpace(string(line)))
		if err == nil {
			result.Numbers[len(result.Numbers)-1] = append(result.Numbers[len(result.Numbers)-1], num)
		} else {
			result.Numbers = append(result.Numbers, []int{})
		}
	}

	for _, val := range lines[len(lines)-2] {
		if val == ' ' {
			continue
		}
		result.Operations = append(result.Operations, OperationFromRune(val))
	}

	return result
}

func applyOperation(a int, b int, op Operation) int {
	switch op {
	case OpSum:
		return a + b
	case OpMul:
		return a * b
	default:
		panic("unknown operation")
	}
}

func Part1(input Input) int {
	result := 0
	for i, op := range input.Operations {
		loc := 0
		for j := 0; j < len(input.Numbers); j++ {
			if j == 0 {
				loc = input.Numbers[j][i]
			} else {
				loc = applyOperation(loc, input.Numbers[j][i], op)
			}
		}
		result += loc
	}
	return result
}

func Part2(input Input) int {
	result := 0
	for i, nums := range input.Numbers {
		loc := 0
		for j, num := range nums {
			if j == 0 {
				loc = num
			} else {
				loc = applyOperation(loc, nums[j], input.Operations[i])
			}
		}
		result += loc
	}
	return result
}
