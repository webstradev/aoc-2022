package day1

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type CaloriesRecord []int

type ElfRecords struct {
	Path    string
	Records []CaloriesRecord
}

func NewElfRecord(path string) *ElfRecords {
	return &ElfRecords{Path: path}
}

func (e *ElfRecords) ReadRecordsFromInputFile() {
	// Initialize AllRecords Object
	e.Records = []CaloriesRecord{}

	// Read the file
	f, err := os.Open(e.Path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create a new scanner to loop over lines
	fs := bufio.NewScanner(f)

	// Create first record
	var record CaloriesRecord

	// Scan over each line in the file
	for fs.Scan() {
		// Ignore whitespace/ empty lines (these are used to seperate elfs)
		if fs.Text() == " " || fs.Text() == "\n" || fs.Text() == "" {
			// If there are entries in the record, add it to the records
			if len(record) > 0 {
				e.Records = append(e.Records, record)
			}

			// Reset Record
			record = CaloriesRecord{}

			// Go to next line because this elf record is complete
			continue
		}

		// Convert the string to an int and panic if that fails
		calorieCount, err := strconv.Atoi(fs.Text())
		if err != nil {
			panic(err)
		}

		// Add the calorie count to the record
		record = append(record, calorieCount)
	}

	// Add the last record to the records (EOF)
	if len(record) > 0 {
		e.Records = append(e.Records, record)
	}
}

func (e *ElfRecords) FindCalorieCountForMaxElf() int {
	if len(e.Records) == 0 {
		panic(errors.New("no records found"))
	}

	maxCount := 0

	for _, record := range e.Records {
		count := 0

		for _, calorieCount := range record {
			count += calorieCount
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func (e *ElfRecords) FindCalorieCountForTopThreeElves() int {
	if len(e.Records) == 0 {
		panic(errors.New("no records found"))
	}

	highest := 0
	secondHighest := 0
	thirdHighest := 0

	for _, record := range e.Records {
		count := 0

		for _, calorieCount := range record {
			count += calorieCount
		}

		if count > thirdHighest {
			if count > secondHighest {
				if count > highest {
					thirdHighest = secondHighest
					secondHighest = highest
					highest = count
				} else {
					thirdHighest = secondHighest
					secondHighest = count
				}
			} else {
				thirdHighest = count
			}
		}
	}

	return highest + secondHighest + thirdHighest
}
