package day3

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type RuckSack struct {
	Items                []string
	SharedItem           string
	PriorityOfSharedItem int
}

var priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (r *RuckSack) FindSharedItem() {
	// Split rucksack into two compartments in the middle
	firstComp := r.Items[0 : len(r.Items)/2]
	secondComp := r.Items[len(r.Items)/2:]

	// Loop over the first compartment
	for _, item := range firstComp {
		// Check if the item is in the second compartment
		for _, item2 := range secondComp {
			if item == item2 {
				r.SharedItem = item
				r.PriorityOfSharedItem = strings.Index(priorities, item) + 1
				// Stop here as we know there is only one
				break
			}
		}
	}

}

type Trip struct {
	Path           string
	Rucksacks      []RuckSack
	PriorityTotal  int
	Badges         []string
	PriorityTotal2 int
}

func NewTrip(path string) *Trip {
	return &Trip{Path: path, Rucksacks: []RuckSack{}}
}

func (t *Trip) ComputeTotalPriorityFromFile() {
	// Read the file
	f, err := os.Open(t.Path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create a new scanner to loop over lines
	fs := bufio.NewScanner(f)

	// Track total priority
	totalPriority := 0

	// Scan over each line in the file
	for fs.Scan() {
		// Ignore whitespace/ empty lines
		if fs.Text() == " " || fs.Text() == "\n" || fs.Text() == "" {
			// Go to next line
			continue
		}

		// Split the line into two strings
		r := RuckSack{Items: strings.Split(fs.Text(), "")}
		r.FindSharedItem()

		// Update total priority
		totalPriority += r.PriorityOfSharedItem

		// Add the rucksack to the trip
		t.Rucksacks = append(t.Rucksacks, r)
	}

	t.PriorityTotal = totalPriority
}

func (t *Trip) ComputeSecondPriorityFromRucksacks() {
	if len(t.Rucksacks) == 0 {
		panic(errors.New("no rucksacks to compute priority from"))
	}

	for i := 0; i < len(t.Rucksacks); i += 3 {
	Rucksack:
		for _, item := range t.Rucksacks[i].Items {
			for _, item2 := range t.Rucksacks[i+1].Items {
				if item == item2 {
					for _, item3 := range t.Rucksacks[i+2].Items {
						if item == item2 && item2 == item3 {
							t.Badges = append(t.Badges, item)
							t.PriorityTotal2 += strings.Index(priorities, item) + 1
							break Rucksack
						}
					}
				}
			}
		}

	}

}
