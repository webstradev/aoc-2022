package day9

import (
	"bufio"
	"os"
)

type Solver struct {
	Path    string
	List    []any
	Result1 int
	Result2 int
}

func NewSolver(path string) *Solver {
	solver := &Solver{
		Path: path,
		List: []any{},
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

	// Scan over each line in the file
	for fs.Scan() {
		text := fs.Text()
		// Ignore whitespace/ empty lines
		if text == " " || text == "\n" || text == "" {
			// Go to next line
			continue
		}

		s.List = append(s.List, text)
	}
}

func (s *Solver) SolvePart1() {

}

func (s *Solver) SolvePart2() {

}
