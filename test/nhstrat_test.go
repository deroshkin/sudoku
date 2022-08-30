package sudoku_test

import (
	"testing"

	"github.com/deroshkin/sudoku/pkg/nhstrats"
	"github.com/deroshkin/sudoku/pkg/solver"
	"github.com/deroshkin/sudoku/util"
)

func TestHiddenSingles(t *testing.T) {
	// No hidden singles
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.HiddenSingles, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden single 1 in row 1 (r1c2)
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.HiddenSingles, true,
		map[util.Cell][]uint8{{R: 0, C: 1}: {1}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden single 2 in box 5 (r4c5)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 6, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.HiddenSingles, true,
		map[util.Cell][]uint8{{R: 3, C: 4}: {2}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Hidden single 7 in column 8 (r8c8)
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 2, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.HiddenSingles, true,
		map[util.Cell][]uint8{{R: 7, C: 7}: {7}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}

func TestHiddenSinglesSolve(t *testing.T) {
	board := [][]uint8{{0, 0, 0, 0, 0, 4, 0, 2, 8},
		{4, 0, 6, 0, 0, 0, 0, 0, 5},
		{1, 0, 0, 0, 3, 0, 6, 0, 0},
		{0, 0, 0, 3, 0, 1, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 1, 4, 0},
		{0, 0, 0, 7, 0, 9, 0, 0, 0},
		{0, 0, 2, 0, 1, 0, 0, 0, 3},
		{9, 0, 0, 0, 0, 0, 5, 0, 7},
		{6, 7, 0, 4, 0, 0, 0, 0, 0}}
	res1, msg1 := util.SolveTester(board, [][]uint8{}, []solver.Strategy{},
		false,
	)
	if !res1 {
		t.Fatalf(msg1)
	}
	boardSolved := [][]uint8{{7, 3, 5, 1, 6, 4, 9, 2, 8},
		{4, 2, 6, 9, 7, 8, 3, 1, 5},
		{1, 9, 8, 5, 3, 2, 6, 7, 4},
		{2, 4, 9, 3, 8, 1, 7, 5, 6},
		{3, 8, 7, 2, 5, 6, 1, 4, 9},
		{5, 6, 1, 7, 4, 9, 8, 3, 2},
		{8, 5, 2, 6, 1, 7, 4, 9, 3},
		{9, 1, 4, 8, 2, 3, 5, 6, 7},
		{6, 7, 3, 4, 9, 5, 2, 8, 1}}
	res2, msg2 := util.SolveTester(board, boardSolved,
		[]solver.Strategy{nhstrats.HiddenSingles}, true,
	)
	if !res2 {
		t.Fatalf(msg2)
	}
}

func TestHiddenSinglesLog(t *testing.T) {
	// Hidden single 1 in row 1 (r1c2)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}

	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.HiddenSingles,
		"Found a hidden single 1 in row 1: r1c2, removing all other candidates from the cell\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden single 7 in column 8 (r8c8)
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 2, 0}}

	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.HiddenSingles,
		"Found a hidden single 7 in column 8: r8c8, removing all other candidates from the cell\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden single 2 in box 5 (r4c5)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 6, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.HiddenSingles,
		"Found a hidden single 2 in box 5: r4c5, removing all other candidates from the cell\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}
}

func TestNakedPairs(t *testing.T) {
	// Naked 12 pair in row 3 (r3c2, r3c3)
	board1 := [][]uint8{{3, 4, 5, 0, 0, 0, 0, 0, 0},
		{6, 7, 8, 0, 0, 0, 0, 0, 0},
		{9, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.NakedPairs, true,
		map[util.Cell][]uint8{{R: 2, C: 1}: {1, 2}, {R: 2, C: 2}: {1, 2}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 2, C: 0}: {1, 2}, {R: 2, C: 3}: {1, 2},
			{R: 2, C: 4}: {1, 2}, {R: 2, C: 5}: {1, 2}, {R: 2, C: 6}: {1, 2},
			{R: 2, C: 7}: {1, 2}, {R: 2, C: 8}: {1, 2}},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Naked 48 pair in column 5 (r1c5, r4c5)
	board2 := [][]uint8{{5, 3, 2, 1, 0, 7, 9, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 3, 2, 0, 0, 0},
		{0, 0, 0, 7, 6, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.NakedPairs, true,
		map[util.Cell][]uint8{{R: 0, C: 4}: {4, 8}, {R: 3, C: 4}: {4, 8}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 1, C: 4}: {4, 8}, {R: 2, C: 4}: {4, 8},
			{R: 4, C: 4}: {4, 8}, {R: 5, C: 4}: {4, 8}, {R: 6, C: 4}: {4, 8},
			{R: 7, C: 4}: {4, 8}, {R: 8, C: 4}: {4, 8}},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Naked 89 pair in box 9 (r7c8, r7c9)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.NakedPairs, true,
		map[util.Cell][]uint8{{R: 6, C: 7}: {8, 9}, {R: 6, C: 8}: {8, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 6, C: 6}: {8, 9}, {R: 7, C: 6}: {8, 9},
			{R: 7, C: 7}: {8, 9}, {R: 7, C: 8}: {8, 9}, {R: 8, C: 6}: {8, 9},
			{R: 8, C: 7}: {8, 9}, {R: 8, C: 8}: {8, 9}},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// No naked pairs
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 0, 4, 5, 6, 7, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.NakedPairs, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}

func TestNakedPairsLog(t *testing.T) {
	// Naked 12 pair in row 3 (r3c2, r3c3)
	board1 := [][]uint8{{3, 4, 5, 0, 0, 0, 0, 0, 0},
		{6, 7, 8, 0, 0, 0, 0, 0, 0},
		{9, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.NakedPairs,
		"Found a naked pair [1 2] in row 3 ( \nr3c2 \nr3c3 \n), removing the values from other cells in the row\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Naked 48 pair in column 5 (r1c5, r4c5)
	board2 := [][]uint8{{5, 3, 2, 1, 0, 7, 9, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 3, 2, 0, 0, 0},
		{0, 0, 0, 7, 6, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.NakedPairs,
		"Found a naked pair [4 8] in column 5 ( \nr1c5 \nr4c5 \n), removing the values from other cells in the column\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Naked 89 pair in box 9 (r7c8, r7c9)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.NakedPairs,
		"Found a naked pair [8 9] in box 9 ( \nr7c8 \nr7c9 \n), removing the values from other cells in the box\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}
}

func TestNakedTriples(t *testing.T) {
	// Naked 129 triple in row 3 (r3c1, r3c2, r3c3)
	board1 := [][]uint8{{3, 4, 5, 0, 0, 0, 0, 0, 0},
		{6, 7, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.NakedTriples, true,
		map[util.Cell][]uint8{{R: 2, C: 0}: {1, 2, 9}, {R: 2, C: 1}: {1, 2, 9},
			{R: 2, C: 2}: {1, 2, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 2, C: 3}: {1, 2, 9}, {R: 2, C: 4}: {1, 2, 9},
			{R: 2, C: 5}: {1, 2, 9}, {R: 2, C: 6}: {1, 2, 9},
			{R: 2, C: 7}: {1, 2, 9}, {R: 2, C: 8}: {1, 2, 9}},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Naked 468 triple in column 5 (r1c5, r4c5, r6c5)
	board2 := [][]uint8{{5, 3, 2, 1, 0, 7, 9, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 3, 2, 0, 0, 0},
		{0, 0, 0, 7, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.NakedTriples, true,
		map[util.Cell][]uint8{{R: 0, C: 4}: {4, 6, 8}, {R: 3, C: 4}: {4, 6, 8},
			{R: 5, C: 4}: {4, 6, 8}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 1, C: 4}: {4, 6, 8}, {R: 2, C: 4}: {4, 6, 8},
			{R: 4, C: 4}: {4, 6, 8}, {R: 6, C: 4}: {4, 6, 8},
			{R: 7, C: 4}: {4, 6, 8}, {R: 8, C: 4}: {4, 6, 8}},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Naked 789 triple in box 9 (r7c7, r7c8, r7c9)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.NakedTriples, true,
		map[util.Cell][]uint8{{R: 6, C: 6}: {7, 8, 9}, {R: 6, C: 7}: {7, 8, 9},
			{R: 6, C: 8}: {7, 8, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 7, C: 6}: {7, 8, 9}, {R: 7, C: 7}: {7, 8, 9},
			{R: 7, C: 8}: {7, 8, 9}, {R: 8, C: 6}: {7, 8, 9},
			{R: 8, C: 7}: {7, 8, 9}, {R: 8, C: 8}: {7, 8, 9}},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Naked 789 triple in box 9 (r7c7, r7c8, r7c9) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.NakedTriples, true,
		map[util.Cell][]uint8{{R: 6, C: 6}: {8, 9}, {R: 6, C: 7}: {7, 9},
			{R: 6, C: 8}: {7, 8}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 7, C: 6}: {7, 8, 9}, {R: 7, C: 7}: {7, 8, 9},
			{R: 7, C: 8}: {7, 8, 9}, {R: 8, C: 6}: {7, 8, 9},
			{R: 8, C: 7}: {7, 8, 9}, {R: 8, C: 8}: {7, 8, 9}},
	)
	if !res4 {
		t.Fatalf(msg4)
	}

	// No naked triples
	board5 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 0, 4, 5, 6, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res5, msg5 := util.CandTester(board5, nhstrats.NakedTriples, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res5 {
		t.Fatalf(msg5)
	}
}

func TestNakedTriplesLog(t *testing.T) {
	// Naked 129 triple in row 3 (r3c1, r3c2, r3c3)
	board1 := [][]uint8{{3, 4, 5, 0, 0, 0, 0, 0, 0},
		{6, 7, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.NakedTriples,
		"Found a naked triple [1 2 9] in row 3 ( \nr3c1 \nr3c2 \nr3c3 \n), removing the values from other cells in the row\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Naked 468 triple in column 5 (r1c5, r4c5, r6c5)
	board2 := [][]uint8{{5, 3, 2, 1, 0, 7, 9, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 3, 2, 0, 0, 0},
		{0, 0, 0, 7, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.NakedTriples,
		"Found a naked triple [4 6 8] in column 5 ( \nr1c5 \nr4c5 \nr6c5 \n), removing the values from other cells in the column\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Naked 789 triple in box 9 (r7c7, r7c8, r7c9)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.NakedTriples,
		"Found a naked triple [7 8 9] in box 9 ( \nr7c7 \nr7c8 \nr7c9 \n), removing the values from other cells in the box\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Naked 789 triple in box 9 (r7c7, r7c8, r7c9) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.LogTester(board4, []solver.Strategy{}, nhstrats.NakedTriples,
		"Found a naked triple [7 8 9] in box 9 ( \nr7c7 \nr7c8 \nr7c9 \n), removing the values from other cells in the box\n",
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}

func TestNakedPairsTriplesSolve(t *testing.T) {
	board := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 9, 0, 0, 5, 0, 0},
		{5, 6, 0, 3, 1, 0, 0, 9, 0},
		{1, 0, 0, 6, 0, 0, 0, 2, 8},
		{0, 0, 4, 0, 0, 0, 7, 0, 0},
		{2, 7, 0, 0, 0, 4, 0, 0, 3},
		{0, 4, 0, 0, 6, 8, 0, 3, 5},
		{0, 0, 2, 0, 0, 5, 9, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	boardSolved := [][]uint8{{9, 2, 8, 5, 4, 7, 3, 1, 6},
		{4, 3, 1, 9, 8, 6, 5, 7, 2},
		{5, 6, 7, 3, 1, 2, 8, 9, 4},
		{1, 9, 5, 6, 7, 3, 4, 2, 8},
		{3, 8, 4, 2, 5, 1, 7, 6, 9},
		{2, 7, 6, 8, 9, 4, 1, 5, 3},
		{7, 4, 9, 1, 6, 8, 2, 3, 5},
		{6, 1, 2, 4, 3, 5, 9, 8, 7},
		{8, 5, 3, 7, 2, 9, 6, 4, 1}}
	res, msg := util.SolveTester(board, boardSolved,
		[]solver.Strategy{nhstrats.HiddenSingles, nhstrats.NakedPairs,
			nhstrats.NakedTriples},
		true,
	)
	if !res {
		t.Fatalf(msg)
	}
}

func TestNakedQuads(t *testing.T) {
	// Naked 6789 quad in row 7 (r7c1, r7c6, r7c7, r7c9)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 1, 0, 0, 2},
		{5, 0, 0, 0, 0, 2, 0, 0, 0},
		{4, 0, 0, 0, 0, 3, 1, 0, 0},
		{0, 0, 0, 0, 0, 4, 2, 0, 0},
		{0, 0, 0, 0, 0, 5, 0, 0, 1},
		{3, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 2, 0, 0, 0, 3, 4, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.NakedQuads, true,
		map[util.Cell][]uint8{{R: 6, C: 0}: {6, 7, 8, 9},
			{R: 6, C: 5}: {6, 7, 8, 9}, {R: 6, C: 6}: {6, 7, 8, 9},
			{R: 6, C: 8}: {6, 7, 8, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 6, C: 1}: {6, 7, 8, 9},
			{R: 6, C: 2}: {6, 7, 8, 9}, {R: 6, C: 3}: {6, 7, 8, 9},
			{R: 6, C: 4}: {6, 7, 8, 9}, {R: 6, C: 7}: {6, 7, 8, 9}},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Naked 1258 quad in column 1 (r1c1, r2c1, r3c1, r8c1)
	board2 := [][]uint8{{0, 3, 4, 0, 0, 0, 0, 0, 0},
		{0, 6, 7, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 3, 4, 6, 7, 9, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.NakedQuads, true,
		map[util.Cell][]uint8{{R: 0, C: 0}: {1, 2, 5, 8},
			{R: 1, C: 0}: {1, 2, 5, 8}, {R: 2, C: 0}: {1, 2, 5, 8},
			{R: 7, C: 0}: {1, 2, 5, 8}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 3, C: 0}: {1, 2, 5, 8},
			{R: 4, C: 0}: {1, 2, 5, 8}, {R: 5, C: 0}: {1, 2, 5, 8},
			{R: 6, C: 0}: {1, 2, 5, 8}, {R: 8, C: 0}: {1, 2, 5, 8}},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Naked 3479 quad in box 2 (r1c4, r2c4, r2c6, r3c4)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 5, 6, 0, 0, 0, 8, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 6, 0, 0, 0, 0, 0},
		{0, 0, 0, 8, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.NakedQuads, true,
		map[util.Cell][]uint8{{R: 0, C: 3}: {3, 4, 7, 9},
			{R: 1, C: 3}: {3, 4, 7, 9}, {R: 1, C: 5}: {3, 4, 7, 9},
			{R: 2, C: 3}: {3, 4, 7, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 0, C: 4}: {3, 4, 7, 9},
			{R: 0, C: 5}: {3, 4, 7, 9}, {R: 1, C: 4}: {3, 4, 7, 9},
			{R: 2, C: 4}: {3, 4, 7, 9}, {R: 2, C: 5}: {3, 4, 7, 9}},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Naked 1234 quad in row 1 (r1c1, r1c3, r1c5, r1c8) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{3, 5, 0, 1, 0, 0, 2, 8, 4},
		{9, 0, 0, 0, 5, 8, 6, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 5, 0},
		{7, 0, 0, 0, 0, 0, 0, 9, 0},
		{8, 0, 0, 0, 7, 0, 0, 0, 0},
		{4, 0, 6, 0, 9, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 0, 6, 0, 0, 7, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.NakedQuads, true,
		map[util.Cell][]uint8{{R: 0, C: 0}: {1, 2},
			{R: 0, C: 2}: {1, 2, 4}, {R: 0, C: 4}: {2, 3, 4},
			{R: 0, C: 7}: {1, 3}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{{R: 0, C: 1}: {1, 2, 3, 4},
			{R: 0, C: 3}: {1, 2, 3, 4}, {R: 0, C: 5}: {1, 2, 3, 4},
			{R: 0, C: 6}: {1, 2, 3, 4}, {R: 0, C: 8}: {1, 2, 3, 4}},
	)
	if !res4 {
		t.Fatalf(msg4)
	}

	// No naked quads
	board5 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res5, msg5 := util.CandTester(board5, nhstrats.NakedQuads, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res5 {
		t.Fatalf(msg5)
	}
}

func TestNakedQuadsLog(t *testing.T) {
	// Naked 6789 quad in row 7 (r7c1, r7c6, r7c7, r7c9)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 1, 0, 0, 2},
		{5, 0, 0, 0, 0, 2, 0, 0, 0},
		{4, 0, 0, 0, 0, 3, 1, 0, 0},
		{0, 0, 0, 0, 0, 4, 2, 0, 0},
		{0, 0, 0, 0, 0, 5, 0, 0, 1},
		{3, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 2, 0, 0, 0, 3, 4, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 0}}
	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.NakedQuads,
		"Found a naked quad [6 7 8 9] in row 7 ( \nr7c1 \nr7c6 \nr7c7 \nr7c9 \n), removing the values from other cells in the row\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Naked 1258 quad in column 1 (r1c1, r2c1, r3c1, r8c1)
	board2 := [][]uint8{{0, 3, 4, 0, 0, 0, 0, 0, 0},
		{0, 6, 7, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 3, 4, 6, 7, 9, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.NakedQuads,
		"Found a naked quad [1 2 5 8] in column 1 ( \nr1c1 \nr2c1 \nr3c1 \nr8c1 \n), removing the values from other cells in the column\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Naked 3479 quad in box 2 (r1c4, r2c4, r2c6, r3c4)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 5, 6, 0, 0, 0, 8, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 6, 0, 0, 0, 0, 0},
		{0, 0, 0, 8, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.NakedQuads,
		"Found a naked quad [3 4 7 9] in box 2 ( \nr1c4 \nr2c4 \nr2c6 \nr3c4 \n), removing the values from other cells in the box\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Naked 1234 quad in row 1 (r1c1, r1c3, r1c5, r1c8) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{3, 5, 0, 1, 0, 0, 2, 8, 4},
		{9, 0, 0, 0, 5, 8, 6, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 5, 0},
		{7, 0, 0, 0, 0, 0, 0, 9, 0},
		{8, 0, 0, 0, 7, 0, 0, 0, 0},
		{4, 0, 6, 0, 9, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 0, 6, 0, 0, 7, 0}}
	res4, msg4 := util.LogTester(board4, []solver.Strategy{}, nhstrats.NakedQuads,
		"Found a naked quad [1 2 3 4] in row 1 ( \nr1c1 \nr1c3 \nr1c5 \nr1c8 \n), removing the values from other cells in the row\n",
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}

func TestHiddenPairs(t *testing.T) {
	// Hidden 12 pair in row 1 (r1c1, r1c2)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.HiddenPairs, true,
		map[util.Cell][]uint8{{R: 0, C: 0}: {1, 2}, {R: 0, C: 1}: {1, 2}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden 47 pair in column 5 (r3c5, r4c5)
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 4, 7, 0},
		{4, 7, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 4, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 0, 0, 7, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.HiddenPairs, true,
		map[util.Cell][]uint8{{R: 2, C: 4}: {4, 7}, {R: 3, C: 4}: {4, 7}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden 56 pair in box 3 (r1c7, r2c8)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 2, 0, 0},
		{5, 6, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.HiddenPairs, true,
		map[util.Cell][]uint8{{R: 0, C: 6}: {5, 6}, {R: 1, C: 7}: {5, 6}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// No hidden pairs
	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.HiddenPairs, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}

func TestHiddenPairsLog(t *testing.T) {
	// Hidden 12 pair in row 1 (r1c1, r1c2)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.HiddenPairs,
		"Found a hidden pair [1 2] in row 1 ( \nr1c1 \nr1c2 \n) removing other values from these cells\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden 47 pair in column 5 (r3c5, r4c5)
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 4, 7, 0},
		{4, 7, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 4, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 0, 0, 7, 0, 0, 0, 0, 0}}
	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.HiddenPairs,
		"Found a hidden pair [4 7] in column 5 ( \nr3c5 \nr4c5 \n) removing other values from these cells\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden 56 pair in box 3 (r1c7, r2c8)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 2, 0, 0},
		{5, 6, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.HiddenPairs,
		"Found a hidden pair [5 6] in box 3 ( \nr1c7 \nr2c8 \n) removing other values from these cells\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}
}

func TestHiddenTriples(t *testing.T) {
	// Hidden 146 triple in row 5 (r5c4, r5c5, r5c6)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 4, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 6, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.HiddenTriples, true,
		map[util.Cell][]uint8{{R: 4, C: 3}: {1, 4, 6}, {R: 4, C: 4}: {1, 4, 6},
			{R: 4, C: 5}: {1, 4, 6}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden 478 triple in column 3 (r4c3, r5c3, r7c3)
	board2 := [][]uint8{{4, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 7, 8},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 7, 8, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.HiddenTriples, true,
		map[util.Cell][]uint8{{R: 3, C: 2}: {4, 7, 8}, {R: 4, C: 2}: {4, 7, 8},
			{R: 6, C: 2}: {4, 7, 8}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden 259 triple in box 5 (r4c4, r4c5, c5c4)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{2, 5, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.HiddenTriples, true,
		map[util.Cell][]uint8{{R: 3, C: 3}: {2, 5, 9}, {R: 3, C: 4}: {2, 5, 9},
			{R: 4, C: 3}: {2, 5, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Hidden 259 triple in box 5 (r4c4, r4c5, r5c4) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{9, 0, 0, 0, 4, 0, 0, 0, 0},
		{2, 5, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.HiddenTriples, true,
		map[util.Cell][]uint8{{R: 3, C: 3}: {2, 5, 9}, {R: 3, C: 4}: {2, 9},
			{R: 4, C: 3}: {2, 5}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res4 {
		t.Fatalf(msg4)
	}

	// No hidden triples
	board5 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res5, msg5 := util.CandTester(board5, nhstrats.HiddenTriples, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res5 {
		t.Fatalf(msg5)
	}
}

func TestHiddenTriplesLog(t *testing.T) {
	// Hidden 146 triple in row 5 (r5c4, r5c5, r5c6)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 4, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 6, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.HiddenTriples,
		"Found a hidden triple [1 4 6] in row 5 ( \nr5c4 \nr5c5 \nr5c6 \n) removing other values from these cells\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden 478 triple in column 3 (r4c3, r5c3, r7c3)
	board2 := [][]uint8{{4, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 7, 8},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 7, 8, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0}}
	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.HiddenTriples,
		"Found a hidden triple [4 7 8] in column 3 ( \nr4c3 \nr5c3 \nr7c3 \n) removing other values from these cells\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden 259 triple in box 5 (r4c4, r4c5, c5c4)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{2, 5, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.HiddenTriples,
		"Found a hidden triple [2 5 9] in box 5 ( \nr4c4 \nr4c5 \nr5c4 \n) removing other values from these cells\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Hidden 259 triple in box 5 (r4c4, r4c5, r5c4) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{9, 0, 0, 0, 4, 0, 0, 0, 0},
		{2, 5, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.LogTester(board4, []solver.Strategy{}, nhstrats.HiddenTriples,
		"Found a hidden triple [2 5 9] in box 5 ( \nr4c4 \nr4c5 \nr5c4 \n) removing other values from these cells\n",
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}

func TestHiddenTripleSolve(t *testing.T) {
	board := [][]uint8{{3, 0, 0, 0, 0, 0, 0, 0, 0},
		{9, 7, 0, 0, 1, 0, 0, 0, 0},
		{6, 0, 0, 5, 8, 3, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 9, 0, 0},
		{5, 0, 0, 6, 2, 1, 0, 0, 3},
		{0, 0, 8, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 4, 3, 5, 0, 0, 2},
		{0, 0, 0, 0, 9, 0, 0, 5, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 1}}
	answer := [][]uint8{{3, 8, 1, 9, 7, 6, 5, 2, 4},
		{9, 7, 5, 2, 1, 4, 6, 3, 8},
		{6, 4, 2, 5, 8, 3, 1, 7, 9},
		{2, 6, 4, 3, 5, 8, 9, 1, 7},
		{5, 9, 7, 6, 2, 1, 4, 8, 3},
		{1, 3, 8, 7, 4, 9, 2, 6, 5},
		{8, 1, 6, 4, 3, 5, 7, 9, 2},
		{4, 2, 3, 1, 9, 7, 8, 5, 6},
		{7, 5, 9, 8, 6, 2, 3, 4, 1}}
	res, msg := util.SolveTester(board, answer,
		[]solver.Strategy{nhstrats.HiddenSingles, nhstrats.NakedTriples, nhstrats.HiddenTriples},
		true,
	)
	if !res {
		t.Fatalf(msg)
	}
}

func TestHiddenQuads(t *testing.T) {
	// Hidden 1234 quad in row 2 (r2c1, r2c2, r2c3, r2c4)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 0, 0, 0, 5, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 2, 3},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.CandTester(board1, nhstrats.HiddenQuads, true,
		map[util.Cell][]uint8{{R: 1, C: 0}: {1, 2, 3, 4},
			{R: 1, C: 1}: {1, 2, 3, 4}, {R: 1, C: 2}: {1, 2, 3, 4},
			{R: 1, C: 3}: {1, 2, 3, 4}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden 2789 quad in column 5 (r1c5, r2c5, r5c5, r6c5)
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 7, 0, 9, 0, 0, 0},
		{0, 0, 0, 8, 0, 0, 0, 0, 0}}
	res2, msg2 := util.CandTester(board2, nhstrats.HiddenQuads, true,
		map[util.Cell][]uint8{{R: 0, C: 4}: {2, 7, 8, 9},
			{R: 1, C: 4}: {2, 7, 8, 9}, {R: 4, C: 4}: {2, 7, 8, 9},
			{R: 5, C: 4}: {2, 7, 8, 9}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden 1456 quad in box 6 (r4c7, r4c8, r6c7, r6c8)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 4, 0, 5, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.CandTester(board3, nhstrats.HiddenQuads, true,
		map[util.Cell][]uint8{{R: 3, C: 6}: {1, 4, 5, 6},
			{R: 3, C: 7}: {1, 4, 5, 6}, {R: 5, C: 6}: {1, 4, 5, 6},
			{R: 5, C: 7}: {1, 4, 5, 6}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Hidden 1235 quad in column 3 (r1c3, r2c3, r5c3, r8c3) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 1, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 3, 0, 0, 2, 0, 0, 5, 0},
		{2, 0, 0, 1, 3, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.CandTester(board4, nhstrats.HiddenQuads, true,
		map[util.Cell][]uint8{{R: 0, C: 2}: {3, 5},
			{R: 1, C: 2}: {1, 2, 3}, {R: 4, C: 2}: {2, 5},
			{R: 7, C: 2}: {1, 3, 5}},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res4 {
		t.Fatalf(msg4)
	}

	board5 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res5, msg5 := util.CandTester(board5, nhstrats.HiddenQuads, false,
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
		map[util.Cell][]uint8{},
	)
	if !res5 {
		t.Fatalf(msg5)
	}
}

func TestHiddenQuadsLog(t *testing.T) {
	// Hidden 1234 quad in row 2 (r2c1, r2c2, r2c3, r2c4)
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 0, 0, 0, 5, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 2, 3},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res1, msg1 := util.LogTester(board1, []solver.Strategy{}, nhstrats.HiddenQuads,
		"Found a hidden quad [1 2 3 4] in row 2 ( \nr2c1 \nr2c2 \nr2c3 \nr2c4 \n) removing other values from these cells\n",
	)
	if !res1 {
		t.Fatalf(msg1)
	}

	// Hidden 2789 quad in column 5 (r1c5, r2c5, r5c5, r6c5)
	board2 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 7, 0, 9, 0, 0, 0},
		{0, 0, 0, 8, 0, 0, 0, 0, 0}}
	res2, msg2 := util.LogTester(board2, []solver.Strategy{}, nhstrats.HiddenQuads,
		"Found a hidden quad [2 7 8 9] in column 5 ( \nr1c5 \nr2c5 \nr5c5 \nr6c5 \n) removing other values from these cells\n",
	)
	if !res2 {
		t.Fatalf(msg2)
	}

	// Hidden 1456 quad in box 6 (r4c7, r4c8, r6c7, r6c8)
	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 4, 0, 5, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	res3, msg3 := util.LogTester(board3, []solver.Strategy{}, nhstrats.HiddenQuads,
		"Found a hidden quad [1 4 5 6] in box 6 ( \nr4c7 \nr4c8 \nr6c7 \nr6c8 \n) removing other values from these cells\n",
	)
	if !res3 {
		t.Fatalf(msg3)
	}

	// Hidden 1235 quad in column 3 (r1c3, r2c3, r5c3, r8c3) with partial values
	board4 := [][]uint8{{0, 0, 0, 0, 1, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 3, 0, 0, 2, 0, 0, 5, 0},
		{2, 0, 0, 1, 3, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 0, 0, 0}}
	res4, msg4 := util.LogTester(board4, []solver.Strategy{}, nhstrats.HiddenQuads,
		"Found a hidden quad [1 2 3 5] in column 3 ( \nr1c3 \nr2c3 \nr5c3 \nr8c3 \n) removing other values from these cells\n",
	)
	if !res4 {
		t.Fatalf(msg4)
	}
}
