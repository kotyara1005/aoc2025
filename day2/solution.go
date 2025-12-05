package day2

import (
	"fmt"
	"iter"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval [2]int

func (it Interval) Contains(num int) bool {
	return it[0] <= num && num <= it[1]
}

type Intervals []Interval

func (its Intervals) Contains(num int) bool {
	i := sort.Search(
		len(its),
		func(i int) bool { return its[i][1] >= num },
	)
	// fmt.Println(num, i, len(its), its)
	if i == len(its) {
		return false
	}
	return its[i].Contains(num)
}

func Parse(filename string) Intervals {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	pairs := strings.Split(string(data), ",")

	result := []Interval{}
	for _, pair := range pairs {
		pair = strings.Trim(pair, "\n")
		nums := strings.Split(pair, "-")
		it := Interval{0, 0}
		it[0], err = strconv.Atoi(nums[0])
		if err != nil {
			panic(err.Error())
		}
		it[1], err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err.Error())
		}
		result = append(result, it)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func atoi(val string) int {
	rv, err := strconv.Atoi(val)
	if err != nil {
		panic(err.Error())
	}
	return rv
}

func GetNextSilly(number int, numberOfParts int) int {
	num := strconv.Itoa(number)
	// Check number of digits
	if len(num)%numberOfParts != 0 {
		silly := "1" + strings.Repeat("0", len(num)/numberOfParts)
		return atoi(strings.Repeat(silly, numberOfParts))
	}
	// Spilit
	half := num[:len(num)/numberOfParts]
	silly := atoi(strings.Repeat(half, numberOfParts))
	if silly > number {
		return silly
	}

	n := atoi(string(half))
	n += 1

	half = strconv.Itoa(n)
	silly = atoi(strings.Repeat(half, 2))

	// Make +1 to upper part
	// Double it
	return silly
}

func GetAllSillyNumberInInterval(it Interval, numberOfParts int) iter.Seq[int] {
	return func(yield func(int) bool) {
		cur := it[0] - 1
		cur = GetNextSilly(cur, numberOfParts)
		for it.Contains(cur) {
			if !yield(cur) {
				return
			}
			cur = GetNextSilly(cur, numberOfParts)
		}
	}
}

func Part1(input Intervals) int {
	result := 0
	for _, it := range input {
		for silly := range GetAllSillyNumberInInterval(it, 2) {
			result += silly
		}
	}
	return result
}

func Part2(input Intervals) int {
	fmt.Println(input)
	result := 0
	seen := make([]bool, 10000000000)
	// go from 1 to 100000
	for i := 1; i < 1000000; i++ {
		// repeat from 2 to 10 times
		for r := 2; r < 12; r++ {
			s := strings.Repeat(strconv.Itoa(i), r)
			if len(s) > 10 {
				break
			}
			silly := atoi(s)
			if seen[silly] {
				continue
			}
			seen[silly] = true
			// check that silly is in some interval
			if input.Contains(silly) {
				fmt.Println("Silly: ", silly)
				result += silly
			}
		}
	}
	return result
}
