package main

import (
	"log"

	"github.com/webstradev/aoc-2022/day01"
	"github.com/webstradev/aoc-2022/day02"
	"github.com/webstradev/aoc-2022/day03"
	"github.com/webstradev/aoc-2022/day04"
	"github.com/webstradev/aoc-2022/day05"
)

func main() {
	// Day 1
	elfRecords := day01.NewElfRecord("day01/calorierecords.txt")
	elfRecords.ReadRecordsFromInputFile()
	part1, part2 := elfRecords.Solve()
	log.Println("=== Day 1 ===")
	log.Printf("Top Elf is carrying %d calories", part1)
	log.Printf("Top 3 Elves are carrying %d calories", part2)

	// Day 2
	game := day02.NewGame("./day02/input.txt")
	game.PlayGameFromFile()
	log.Println("=== Day 2 ===")
	log.Printf("Total score is %d", game.TotalScore)
	log.Printf("Total score the real way is %d", game.TotalScore2)

	// Day 3
	trip := day03.NewTrip("./day03/input.txt")
	trip.ComputeTotalPriorityFromFile()
	trip.ComputeSecondPriorityFromRucksacks()
	log.Println("=== Day 3 ===")
	log.Printf("Total priority is %d", trip.PriorityTotal)
	log.Printf("Total priority for badges is %d", trip.PriorityTotal2)

	// Day 4
	solver := day04.NewAssignmentsSolver("./day04/input.txt")
	solver.FindNumberOfUselessAssignments()
	solver.FindNumberOfOverlappingAssignments()
	log.Println("=== Day 4 ===")
	log.Println("Number of useless assignments is", solver.NumberOfUselessAssignments)
	log.Println("Number of overlapping assignments is", solver.NumberOfOverlappingAssignments)

	// Day 5
	craneSolver := day05.NewSolver("./day05/input.txt")
	craneSolver.SolvePart1()
	craneSolver.SolvePart2()
	log.Println("=== Day 5 ===")
	log.Println("Top Crates for crane 9000 are", craneSolver.Result1)
	log.Println("Top Crates for crane 9001 are", craneSolver.Result2)
}
