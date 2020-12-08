package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	pwData := strings.Split(ReadFile("input/day2"), "\n")

	partOneValidPws := 0

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
	}

	fmt.Printf("Part One Valid PWs: %v", partOneValidPws)
}
