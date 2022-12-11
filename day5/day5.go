package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Stack []string

type State struct {
	Stacks []Stack
}

type Assignment struct {
	Amount int
	From   int
	To     int
}

type CraneAssignment struct {
	State        State
	Instructions []Assignment
}

type Solver struct {
	Path            string
	OrignalState    State
	CraneAssignment CraneAssignment
	Result1         string
	Result2         string
}

func NewSolver(path string) *Solver {
	solver := &Solver{
		Path:            path,
		CraneAssignment: CraneAssignment{},
	}

	solver.ReadListFromFile()

	return solver
}

func (s *Solver) ReadListFromFile() {
	// Read the file
	f, err := os.Open(s.Path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create a new scanner to loop over lines
	fs := bufio.NewScanner(f)

	// Track whether we are in the state or assignments section
	stateDone := false

	// Initialize the stacks
	s.CraneAssignment.State.Stacks = []Stack{}

	// Scan over each line in the file
	for fs.Scan() {
		text := fs.Text()

		if fs.Text() == "" {
			stateDone = true
			continue
		}

		if !stateDone {
			cratesOnLine := []string{}
			// One stack is always 3 characters wide

			for i := 0; i < len(text); i += 4 { // we skip 4 characters because there is a space between each stack
				// Get the 3 characters of this stack
				stack := text[i : i+3]

				if stack == "   " {
					cratesOnLine = append(cratesOnLine, " ")
					continue // go to the next stack but stay in the same line
				}

				// Convert the stack to an int
				cratesOnLine = append(cratesOnLine, string(stack[1]))
			}

			// Add the crates to the actual stack
			for i := 0; i < len(cratesOnLine); i++ {
				if len(s.CraneAssignment.State.Stacks) <= i {
					s.CraneAssignment.State.Stacks = append(s.CraneAssignment.State.Stacks, Stack{})
				}

				// Add the crate to the stack
				if cratesOnLine[i] != " " {
					s.CraneAssignment.State.Stacks[i] = append(s.CraneAssignment.State.Stacks[i], string(cratesOnLine[i]))
				}
			}
		}

		if stateDone {
			splits := strings.Split(text, " ")
			amount, err := strconv.Atoi(splits[1])
			if err != nil {
				panic(err)
			}

			from, err := strconv.Atoi(splits[3])
			if err != nil {
				panic(err)
			}

			to, err := strconv.Atoi(splits[5])
			if err != nil {
				panic(err)
			}

			s.CraneAssignment.Instructions = append(s.CraneAssignment.Instructions, Assignment{
				Amount: amount,
				From:   from,
				To:     to,
			})
		}
	}

	s.OrignalState = State{
		Stacks: append([]Stack{}, s.CraneAssignment.State.Stacks...),
	}
}

func (s *Solver) SolvePart1() {
	// Move the crates following instructions
	for _, instruction := range s.CraneAssignment.Instructions {
		for i := 0; i < instruction.Amount; i++ {
			crate := s.CraneAssignment.State.Stacks[instruction.From-1][0]

			// Add the crate to the "to" stack
			s.CraneAssignment.State.Stacks[instruction.To-1] = append([]string{crate}, s.CraneAssignment.State.Stacks[instruction.To-1]...)

			// Remove the crate from the "from" stack
			s.CraneAssignment.State.Stacks[instruction.From-1] = s.CraneAssignment.State.Stacks[instruction.From-1][1:]
		}
	}

	// Find the top crate on each stack
	for _, stack := range s.CraneAssignment.State.Stacks {
		s.Result1 += stack[0]
	}
}

func (s *Solver) SolvePart2() {
	// Reset the state
	s.CraneAssignment.State = s.OrignalState

	// Move the crates following instructions
	for _, instruction := range s.CraneAssignment.Instructions {
		from := append([]string{}, s.CraneAssignment.State.Stacks[instruction.From-1]...)
		to := append([]string{}, s.CraneAssignment.State.Stacks[instruction.To-1]...)
		crates := append([]string{}, from[0:instruction.Amount]...)

		// Remove the crate from the "from" stack
		newFromStack := from[instruction.Amount:]

		newToStack := append(crates, to...)

		// Add the crates to the "to" stack
		s.CraneAssignment.State.Stacks[instruction.To-1] = newToStack

		// Remove the crates from the "from" stack
		s.CraneAssignment.State.Stacks[instruction.From-1] = newFromStack
	}

	// Find the top crate on each stack
	for _, stack := range s.CraneAssignment.State.Stacks {
		s.Result2 += stack[0]
	}
}
