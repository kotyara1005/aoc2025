package day2

import (
	"testing"
)

func TestGetNextSilly(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		number int
		want   int
	}{
		{"Odd digits", 101, 1010},
		{"Even", 1010, 1111},
		{"Even", 1110, 1111},
		{"Even", 1112, 1212},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNextSilly(tt.number, 2); got != tt.want {
				t.Errorf("GetNextSilly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"Test", "test_input", 1227775554},
		{"Real", "input", 29940924880},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(Parse(tt.filename)); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"Test", "test_input", 4174379265},
		{"Real", "input", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(Parse(tt.filename)); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervals_Contains(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"", 15, true},
		{"", -1, false},
		{"", 446449, true},
		{"", 999999999, false},
		{"", 1188511882, true},
	}
	for _, tt := range tests {
		its := Parse("test_input")
		t.Run(tt.name, func(t *testing.T) {
			if got := its.Contains(tt.num); got != tt.want {
				t.Errorf("Intervals.Contains(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
