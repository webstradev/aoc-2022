package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver_SolvePart1(t *testing.T) {
	tests := []struct {
		name string
		s    *Solver
		want string
	}{
		{
			name: "Example from aoc",
			s:    NewSolver("./testdata/input.txt"),
			want: "CMZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SolvePart1()
			assert.Equal(t, tt.want, tt.s.Result1)
		})
	}
}

func TestSolver_SolvePart2(t *testing.T) {
	tests := []struct {
		name string
		s    *Solver
		want string
	}{
		{
			name: "Example from aoc",
			s:    NewSolver("./testdata/input.txt"),
			want: "MCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SolvePart2()
			assert.Equal(t, tt.want, tt.s.Result2)
		})
	}
}
