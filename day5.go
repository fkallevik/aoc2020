package main

import (
	"fmt"
	"sort"
	"strings"
)

type boardingPass struct {
	row    int
	column int
	seatId int
}

func lowerHalf(upper int, lower int) int {
	return upper - rowCount(upper, lower)/2
}

func upperHalf(upper int, lower int) int {
	return lower + rowCount(upper, lower)/2
}

func rowCount(upper int, lower int) int {
	return upper - lower + 1
}

func Day5() {
	input := strings.Split(ReadFile("input/day5"), "\n")

	var boardingPasses []boardingPass

	for _, bpSeatCode := range input {
		upperRow := 127
		lowerRow := 0

		upperCol := 7
		lowerCol := 0

		for _, l := range bpSeatCode {
			switch string(l) {
			case "F":
				upperRow = lowerHalf(upperRow, lowerRow)
				break
			case "B":
				lowerRow = upperHalf(upperRow, lowerRow)
				break
			case "L":
				upperCol = lowerHalf(upperCol, lowerCol)
				break
			case "R":
				lowerCol = upperHalf(upperCol, lowerCol)
				break
			default:
				break
			}
		}

		boardingPasses = append(boardingPasses, boardingPass{
			row:    lowerRow,
			column: lowerCol,
			seatId: lowerRow*8 + lowerCol,
		})
	}

	var seatIds []int

	for _, bp := range boardingPasses {
		seatIds = append(seatIds, bp.seatId)
	}

	sort.Ints(seatIds)

	var mySeatId int

	for i, sid := range seatIds {
		if i > 0 && seatIds[i-1]+1 != sid {
			mySeatId = sid - 1
		}
	}

	fmt.Printf("\nPart One - Highest Seat ID: %v\n", seatIds[len(seatIds)-1])
	fmt.Printf("\nPart Two - My Seat ID: %v\n", mySeatId)
}
