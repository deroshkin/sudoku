package sudoku_test

import (
	"testing"

	"github.com/deroshkin/sudoku/pkg/nhstrats"
	"github.com/deroshkin/sudoku/pkg/solver"
	"github.com/deroshkin/sudoku/util"
)

func TestHiddenSingles(t *testing.T) {
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

	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0}}
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

func TestNakedPairs(t *testing.T) {
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

func TestNakedTriples(t *testing.T) {
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
