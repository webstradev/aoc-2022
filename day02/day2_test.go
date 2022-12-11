package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_PlayGameFromFile(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedScore  int
		expectedScore2 int
	}{
		{
			name:           "Example from aoc",
			path:           "./testdata/input.txt",
			expectedScore:  15,
			expectedScore2: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGame(tt.path)
			g.PlayGameFromFile()
			assert.Equal(t, g.TotalScore, tt.expectedScore)
			assert.Equal(t, g.TotalScore2, tt.expectedScore2)
		})
	}
}
