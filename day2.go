package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getEndPos(p int, s string) int {
	if len(s) < p+1 {
		return len(s)
	}

	return p + 1
}

func Day2() {
	pwData := strings.Split(ReadFile("input/day2"), "\n")

	partOneValidPws := 0
	partTwoValidPws := 0

	for _, line := range pwData {
		// Parse password and policy data.

		splitLine := strings.Split(line, " ")

		policyNumbers := strings.Split(splitLine[0], "-")
		pwLetter := strings.Trim(splitLine[1], ":")
		pw := splitLine[2]

		count := strings.Count(pw, pwLetter)

		policyNumber1, err := strconv.Atoi(policyNumbers[0])

		if err != nil {
			panic(err)
		}

		policyNumber2, err := strconv.Atoi(policyNumbers[1])

		if err != nil {
			panic(err)
		}

		// Validate password for part one.

		min := policyNumber1
		max := policyNumber2

		if count >= min && count <= max {
			partOneValidPws++
		}

		// Validate password for part two.
		pos1 := policyNumber1 - 1
		pos2 := policyNumber2 - 1

		pos1Char := pw[pos1:getEndPos(pos1, pw)]
		pos2Char := pw[pos2:getEndPos(pos2, pw)]

		if (pos1Char == pwLetter && pos2Char != pwLetter) || (pos2Char == pwLetter && pos1Char != pwLetter) {
			partTwoValidPws++
		}
	}

	fmt.Printf("Part One Valid PWs: %v\n", partOneValidPws)
	fmt.Printf("Part Two Valid PWs: %v", partTwoValidPws)
}
