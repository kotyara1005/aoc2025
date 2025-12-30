package day8

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		filename string
		iters int
		want int
	}{
		{"Test", "test_input", 10, 40},
		{"Real", "input", 1000, 96672},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(Parse(tt.filename), tt.iters); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name string
		filename string
		want int
	}{
		{"Test", "test_input", 25272},
		{"Real", "input", 22517595},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(Parse(tt.filename)); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
