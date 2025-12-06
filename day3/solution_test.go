package day3

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"Test", "test_input", 357},
		{"Real", "input", 17445},
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
		{"Test", "test_input", 3121910778619},
		{"Real", "input", 173229689350551},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(Parse(tt.filename)); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
