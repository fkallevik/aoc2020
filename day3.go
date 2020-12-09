package main

import (
	"fmt"
	"strings"
)

const TREE = "#"
const SPACE = "."

func Day3() {
	mapPattern := strings.Split(ReadFile("input/day3"), "\n")

	trees := 0

	x := 0

	for y := 1; y < len(mapPattern); y++ {
		x += 3

		row := mapPattern[y]
		maxPosX := len(row) - 1

		// Shift x value to simulate map pattern repeat.
		if x > maxPosX {
			x -= len(row)
		}

		if x > len(mapPattern)-3 {
			break
		}

		runes := []rune(row)
		objectAtPos := string(runes[x])

		if objectAtPos == TREE {
			trees++
			runes[x] = []rune("X")[0]
		}

		if objectAtPos == SPACE {
			runes[x] = []rune("0")[0]
		}

		fmt.Printf("%v\n", string(runes))
	}

	fmt.Printf("\nPart One Trees: %v\n", trees)
}
