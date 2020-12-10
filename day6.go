package main

import (
	"fmt"
	"strings"
)

type group struct {
	size      int
	questions map[string]int
}

func newGroup() group {
	return group{
		size:      0,
		questions: make(map[string]int),
	}
}

func Day6() {
	input := strings.Split(ReadFile("input/day6"), "\n")

	var groups []group

	group := newGroup()

	for _, person := range input {
		if person == "" {
			groups = append(groups, group)
			group = newGroup()
			continue
		}

		group.size++

		for _, question := range person {
			group.questions[string(question)] += 1
		}
	}

	// Append the last group.
	groups = append(groups, group)

	partOneSum := 0
	partTwoSum := 0

	for _, group := range groups {
		partOneSum += len(group.questions)

		for _, q := range group.questions {
			if q == group.size {
				partTwoSum += 1
			}
		}
	}

	fmt.Printf("\nPart One: Sum of questions for all groups: %v\n", partOneSum)
	fmt.Printf("\nPart Two: Unanimous questions: %v\n", partTwoSum)
}
