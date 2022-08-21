package sudoku_test

import (
	"github.com/deroshkin/sudoku/pkg/interstrats"
	"github.com/deroshkin/sudoku/util"
	"testing"
)

func TestPointingSets(t *testing.T) {
	// Pointing pair of 2s in row 1
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 1, 0},
		{0, 0, 0, 0, 0, 0, 7, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, interstrats.PointingSets, true,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 0, C: 6}: {2}, {R: 0, C: 7}: {2}},
		map[util.Cell][]uint8{{R: 0, C: 0}: {2}, {R: 0, C: 1}: {2}, {R: 0, C: 2}: {2},
			{R: 0, C: 3}: {2}, {R: 0, C: 4}: {2}, {R: 0, C: 5}: {2}},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Pointing pair of 7s in column 3
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{1, 2, 0, 0, 0, 0, 0, 0, 0},
		{5, 4, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, interstrats.PointingSets, true,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 4, C: 2}: {7}, {R: 5, C: 2}: {7}},
		map[util.Cell][]uint8{{R: 0, C: 2}: {7}, {R: 1, C: 2}: {7}, {R: 2, C: 2}: {7},
			{R: 6, C: 2}: {7}, {R: 7, C: 2}: {7}, {R: 8, C: 2}: {7}},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	//Pointing triple of 4s in row 5
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 2, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, interstrats.PointingSets, true,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 4, C: 3}: {4}, {R: 4, C: 4}: {4}, {R: 4, C: 5}: {4}},
		map[util.Cell][]uint8{{R: 4, C: 0}: {4}, {R: 4, C: 1}: {4}, {R: 4, C: 2}: {4},
			{R: 4, C: 6}: {4}, {R: 4, C: 7}: {4}, {R: 4, C: 8}: {4}},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	//Pointing triple of 8s in column 1
	board4 := [][]uint8{{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 5, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.CandTester(board4, interstrats.PointingSets, true,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 0, C: 0}: {8}, {R: 1, C: 0}: {8}, {R: 2, C: 0}: {8}},
		map[util.Cell][]uint8{{R: 3, C: 0}: {8}, {R: 4, C: 0}: {8}, {R: 5, C: 0}: {8},
			{R: 6, C: 0}: {8}, {R: 7, C: 0}: {8}, {R: 8, C: 0}: {8}},
	)
	if !res4 {
		t.Fatalf(msg4)
	}

	//Non-example
	board5 := [][]uint8{{0, 1, 4, 0, 0, 0, 0, 0, 0},
		{0, 2, 6, 0, 0, 0, 0, 0, 0},
		{0, 5, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res5, msg5 := util.CandTester(board5, interstrats.PointingSets, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res5 {
		t.Fatalf(msg5)
	}
}