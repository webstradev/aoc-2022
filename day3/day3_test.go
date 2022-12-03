package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrip_ComputeTotalPriorityFromFile(t *testing.T) {
	tests := []struct {
		name             string
		tr               *Trip
		expectedPriority int
	}{
		{
			name:             "example from aoc",
			tr:               NewTrip("./testdata/input.txt"),
			expectedPriority: 157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.ComputeTotalPriorityFromFile()
			assert.Equal(t, tt.expectedPriority, tt.tr.PriorityTotal)
		})
	}
}
