package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignmentsSolver_FindNumberOfUselessAssignments(t *testing.T) {
	tests := []struct {
		name string
		s    *AssignmentsSolver
		want int
	}{
		{
			name: "Example from aoc",
			s:    NewAssignmentsSolver("./testdata/input.txt"),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.FindNumberOfUselessAssignments()
			assert.Equal(t, tt.want, tt.s.NumberOfUselessAssignments)
		})
	}
}

func TestAssignmentsSolver_FindNumberOfOverlappingAssignments(t *testing.T) {
	tests := []struct {
		name string
		s    *AssignmentsSolver
		want int
	}{
		{
			name: "Example from aoc",
			s:    NewAssignmentsSolver("./testdata/input.txt"),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.FindNumberOfOverlappingAssignments()
			assert.Equal(t, tt.want, tt.s.NumberOfOverlappingAssignments)
		})
	}
}
