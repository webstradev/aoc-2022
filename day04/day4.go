package day04

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	Min int
	Max int
}

type PairedAssignment struct {
	First Assignment
	Last  Assignment
}

type AssignmentsSolver struct {
	Path                           string
	Assignments                    []PairedAssignment
	NumberOfUselessAssignments     int
	NumberOfOverlappingAssignments int
}

func NewAssignmentsSolver(path string) *AssignmentsSolver {
	solver := &AssignmentsSolver{
		Path:        path,
		Assignments: []PairedAssignment{},
	}

	solver.ReadAssignmentsFromFile()

	return solver
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (s *AssignmentsSolver) ReadAssignmentsFromFile() {
	// Read the file
	f, err := os.Open(s.Path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create a new scanner to loop over lines
	fs := bufio.NewScanner(f)

	// Scan over each line in the file
	for fs.Scan() {
		assignment := fs.Text()
		// Ignore whitespace/ empty lines
		if assignment == " " || assignment == "\n" || assignment == "" {
			// Go to next line
			continue
		}

		pairs := strings.Split(assignment, ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")

		s.Assignments = append(s.Assignments, PairedAssignment{
			First: Assignment{
				Min: stringToInt(first[0]),
				Max: stringToInt(first[1]),
			},
			Last: Assignment{
				Min: stringToInt(second[0]),
				Max: stringToInt(second[1]),
			},
		},
		)
	}
}

func (s *AssignmentsSolver) FindNumberOfUselessAssignments() {
	count := 0
	for _, assignment := range s.Assignments {
		firstIsinLast := assignment.First.Min >= assignment.Last.Min && assignment.First.Max <= assignment.Last.Max
		lastIsInFirst := assignment.Last.Min >= assignment.First.Min && assignment.Last.Max <= assignment.First.Max
		if firstIsinLast || lastIsInFirst {
			count++
			continue
		}
	}
	s.NumberOfUselessAssignments = count
}

func (s *AssignmentsSolver) FindNumberOfOverlappingAssignments() {
	count := 0
	for _, assignment := range s.Assignments {
	NextAssignment:
		for i := assignment.First.Min; i <= assignment.First.Max; i++ {
			if i >= assignment.Last.Min && i <= assignment.Last.Max {
				count++
				break NextAssignment
			}
		}
	}
	s.NumberOfOverlappingAssignments = count
}
