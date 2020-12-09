package main

import (
	"fmt"
	"strings"
)

const TREE = "#"
const SPACE = "."

var mapPattern []string

func init() {
	mapPattern = strings.Split(ReadFile("input/day3"), "\n")
}

func calcTreesForSlope(right int, down int) int {
	trees := 0

	x := 0

	for y := down; y < len(mapPattern); y += down {
		x += right

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

	return trees
}

func Day3() {
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	partOneTrees := 0
	partTwoTrees := 1

	for i, slope := range slopes {
		t := calcTreesForSlope(slope[0], slope[1])
		partTwoTrees *= t

		if i == 1 {
			partOneTrees = t
		}
	}

	fmt.Printf("\nPart One - Trees Encountered: %v\n", partOneTrees)
	fmt.Printf("\nPart Two - All Slope Trees Multiplied: %v\n", partTwoTrees)
}
