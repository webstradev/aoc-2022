package main

import (
	"log"

	"github.com/webstradev/aoc-2022/day1"
	"github.com/webstradev/aoc-2022/day2"
	"github.com/webstradev/aoc-2022/day3"
)

func main() {
	// Day 1
	elfRecords := day1.NewElfRecord("day1/calorierecords.txt")
	elfRecords.ReadRecordsFromInputFile()
	part1, part2 := elfRecords.Solve()
	log.Printf("Top Elf is carrying %d calories", part1)
	log.Printf("Top 3 Elves are carrying %d calories", part2)

	// Day 2
	game := day2.NewGame("./day2/input.txt")
	game.PlayGameFromFile()
	log.Printf("Total score is %d", game.TotalScore)
	log.Printf("Total score the real way is %d", game.TotalScore2)

	// Day 3
	trip := day3.NewTrip("./day3/input.txt")
	trip.ComputeTotalPriorityFromFile()
	log.Printf("Total priority is %d", trip.PriorityTotal)
}
