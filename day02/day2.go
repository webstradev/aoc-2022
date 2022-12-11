package day02

import (
	"bufio"
	"os"
	"strings"
)

type Round struct {
	Opponent string
	You      string
	Result   int //puzzle1
	Result2  int //puzzle2
}

type Game struct {
	Path        string
	Strategy    []Round
	TotalScore  int //puzzle1
	TotalScore2 int //puzzle2
}

var mapXtoA = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var scoreForChoice = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var scoreForLoss = 0
var scoreForDraw = 3
var scoreForWin = 6

var scoreFromX = map[string]int{
	"X": scoreForLoss,
	"Y": scoreForDraw,
	"Z": scoreForWin,
}

var mapOutcomeAndOpponentToYou = map[string]map[string]string{
	"X": {
		"A": "C",
		"B": "A",
		"C": "B",
	},
	"Y": {
		"A": "A",
		"B": "B",
		"C": "C",
	},
	"Z": {
		"A": "B",
		"B": "C",
		"C": "A",
	},
}

func (r *Round) PlayRound() {
	you := mapXtoA[r.You]
	// Draw
	if r.Opponent == you {
		r.Result = scoreForDraw + scoreForChoice[you]
		return
	}

	// Win
	if r.Opponent == "A" && you == "B" ||
		r.Opponent == "B" && you == "C" ||
		r.Opponent == "C" && you == "A" {
		r.Result = scoreForWin + scoreForChoice[you]
		return
	}

	// loss
	r.Result = scoreForLoss + scoreForChoice[you]
}

func (r *Round) PlayRound2() {
	r.Result2 = scoreFromX[r.You] + scoreForChoice[mapOutcomeAndOpponentToYou[r.You][r.Opponent]]
}

func NewGame(path string) *Game {
	return &Game{Path: path, Strategy: []Round{}}
}

func (g *Game) PlayGameFromFile() {
	// Read the file
	f, err := os.Open(g.Path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create a new scanner to loop over lines
	fs := bufio.NewScanner(f)

	// Tracker variables
	totalScore := 0
	totalScore2 := 0

	// Scan over each line in the file
	for fs.Scan() {
		// Ignore whitespace/ empty lines
		if fs.Text() == " " || fs.Text() == "\n" || fs.Text() == "" {
			// Go to next line
			continue
		}

		// Split the line into two strings
		split := strings.Split(fs.Text(), " ")

		round := Round{Opponent: split[0], You: split[1]}
		round.PlayRound()
		round.PlayRound2()

		// Update trackers
		totalScore += round.Result
		totalScore2 += round.Result2

		// Add the move to the strategy
		g.Strategy = append(g.Strategy, round)
	}

	g.TotalScore = totalScore
	g.TotalScore2 = totalScore2
}
