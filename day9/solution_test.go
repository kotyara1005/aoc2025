package day9

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name   string
		filename string
		want   int
	}{
		{"Test", "test_input", 50},
		{"Real", "input", 4748826374},
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
		name   string
		filename string
		want   int
	}{
		{"Test", "test_input", 24},
		{"Real", "input", 1554370486},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(Parse(tt.filename)); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
