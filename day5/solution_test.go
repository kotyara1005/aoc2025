package day5

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"Test", "test_input", 3},
		{"Real", "input", 888},
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
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"Test", "test_input", 14},
		{"Real", "input", 344378119285354},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(Parse(tt.filename)); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
