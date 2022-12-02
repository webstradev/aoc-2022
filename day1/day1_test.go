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

func TestElfRecords_FindCalorieCountForMaxElf(t *testing.T) {
	tests := []struct {
		name string
		e    *ElfRecords
		want int
	}{
		{
			name: "Test ReadRecordsFromInputFile",
			e:    NewElfRecord("./testdata/calorierecords.txt"),
			want: 24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.ReadRecordsFromInputFile()
			assert.Equal(t, tt.e.FindCalorieCountForMaxElf(), tt.want)
		})
	}
}

func TestElfRecords_FindCalorieCountForTopThreeElves(t *testing.T) {
	tests := []struct {
		name string
		e    *ElfRecords
		want int
	}{
		{
			name: "Test ReadRecordsFromInputFile",
			e:    NewElfRecord("./testdata/calorierecords.txt"),
			want: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				tt.e.ReadRecordsFromInputFile()
				assert.Equal(t, tt.e.FindCalorieCountForTopThreeElves(), tt.want)
			})
		})
	}
}
