package main

import (
	"log"

	"github.com/webstradev/aoc-2022/day1"
)

func main() {
	// Day 1
	elfRecords := day1.NewElfRecord("day1/calorierecords.txt")
	elfRecords.ReadRecordsFromInputFile()
	part1 := elfRecords.FindCalorieCountForMaxElf()
	part2 := elfRecords.FindCalorieCountForTopThreeElves()
	log.Printf("Top Elf is carrying %d calories", part1)
	log.Printf("Top 3 Elves are carrying %d calories", part2)
}
