package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElfRecords_ReadRecordsFromInputFile(t *testing.T) {
	tests := []struct {
		name string
		e    *ElfRecords
		want []CaloriesRecord
	}{
		{
			name: "Test ReadRecordsFromInputFile",
			e:    NewElfRecord("./testdata/calorierecords.txt"),
			want: []CaloriesRecord{
				{1000, 2000, 3000},
				{4000},
				{5000, 6000},
				{7000, 8000, 9000},
				{10000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.ReadRecordsFromInputFile()

			assert.Equal(t, tt.e.Records, tt.want)
		})
	}
}

func TestElfRecords_Solve(t *testing.T) {
	tests := []struct {
		name  string
		e     *ElfRecords
		want1 int
		want2 int
	}{
		{
			name:  "Test ReadRecordsFromInputFile",
			e:     NewElfRecord("./testdata/calorierecords.txt"),
			want1: 24000,
			want2: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				tt.e.ReadRecordsFromInputFile()
				part1, part2 := tt.e.Solve()
				assert.Equal(t, part1, tt.want1)
				assert.Equal(t, part2, tt.want2)
			})
		})
	}
}
