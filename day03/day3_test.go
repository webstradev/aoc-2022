package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrip_ComputeTotalPriorityFromFile(t *testing.T) {
	tests := []struct {
		name              string
		tr                *Trip
		expectedPriority  int
		expectedPriority2 int
	}{
		{
			name:              "example from aoc",
			tr:                NewTrip("./testdata/input.txt"),
			expectedPriority:  157,
			expectedPriority2: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.ComputeTotalPriorityFromFile()
			tt.tr.ComputeSecondPriorityFromRucksacks()
			assert.Equal(t, tt.expectedPriority, tt.tr.PriorityTotal)
			assert.Equal(t, tt.expectedPriority2, tt.tr.PriorityTotal2)
		})
	}
}
