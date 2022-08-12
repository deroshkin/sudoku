package sudoku_test

import (
	"testing"

	"github.com/deroshkin/sudoku/pkg/nhstrats"
	"github.com/deroshkin/sudoku/pkg/solver"
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
	sol1 := solver.MakeSolver(board1, []solver.Strategy{})
	if nhstrats.HiddenSingles(sol1.Cands) {
		t.Fatalf("Board %v should not have any hidden singles", board1)
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
	sol2 := solver.MakeSolver(board2, []solver.Strategy{})
	sol2.RestrictCands()
	if !nhstrats.HiddenSingles(sol2.Cands) {
		t.Fatalf("Board %v should have a hidden single 1 in row 1 (r1c2)", board2)
	}
	if len(sol2.Cands[0][1]) != 1 {
		t.Fatalf("Hidden single in r1c2 not restricted in %v, still %v candidates", board2, len(sol2.Cands[0][1]))
	}
	if sol2.Cands[0][1][0] != 1 {
		t.Fatalf("Incorrect hidden single in r1c2 of %v, expected 1, got %v", board2, sol2.Cands[0][1][0])
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
	sol3 := solver.MakeSolver(board3, []solver.Strategy{})
	sol3.RestrictCands()
	if !nhstrats.HiddenSingles(sol3.Cands) {
		t.Fatalf("Board %v should have a hidden single 2 in box 5 (r4c5)", board3)
	}
	if len(sol3.Cands[3][4]) != 1 {
		t.Fatalf("Hidden single in r4c5 not restricted in %v, still %v candidates", board3, len(sol3.Cands[3][4]))
	}
	if sol3.Cands[3][4][0] != 2 {
		t.Fatalf("Incorrect hidden single in r4c5 of %v, expected 2, got %v", board3, sol3.Cands[3][4][0])
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
	sol4 := solver.MakeSolver(board4, []solver.Strategy{})
	sol4.RestrictCands()
	if !nhstrats.HiddenSingles(sol4.Cands) {
		t.Fatalf("Board %v should have a hidden single 7 in column 8 (r8c8)", board4)
	}
	if len(sol4.Cands[7][7]) != 1 {
		t.Fatalf("Hidden single in r8c8 not restricted in %v, still %v candidates", board4, len(sol4.Cands[7][7]))
	}
	if sol4.Cands[7][7][0] != 7 {
		t.Fatalf("Incorrect hidden single in r8c8 of %v, expected 7, got %v", board4, sol4.Cands[7][7][0])
	}
}

func TestHiddenSinglesSolve(t *testing.T) {
	board1 := [][]uint8{{0, 0, 0, 0, 0, 4, 0, 2, 8},
		{4, 0, 6, 0, 0, 0, 0, 0, 5},
		{1, 0, 0, 0, 3, 0, 6, 0, 0},
		{0, 0, 0, 3, 0, 1, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 1, 4, 0},
		{0, 0, 0, 7, 0, 9, 0, 0, 0},
		{0, 0, 2, 0, 1, 0, 0, 0, 3},
		{9, 0, 0, 0, 0, 0, 5, 0, 7},
		{6, 7, 0, 4, 0, 0, 0, 0, 0}}
	sol1a := solver.MakeSolver(board1, []solver.Strategy{})
	if sol1a.Solve() {
		t.Fatalf("Board %v should not be solvable without extra strategies", board1)
	}
	sol1b := solver.MakeSolver(board1, []solver.Strategy{nhstrats.HiddenSingles})
	if !sol1b.Solve() {
		t.Fatalf("Board %v should be solvable using only hidden singles", board1)
	}
	board1Solved := [][]uint8{{7, 3, 5, 1, 6, 4, 9, 2, 8},
		{4, 2, 6, 9, 7, 8, 3, 1, 5},
		{1, 9, 8, 5, 3, 2, 6, 7, 4},
		{2, 4, 9, 3, 8, 1, 7, 5, 6},
		{3, 8, 7, 2, 5, 6, 1, 4, 9},
		{5, 6, 1, 7, 4, 9, 8, 3, 2},
		{8, 5, 2, 6, 1, 7, 4, 9, 3},
		{9, 1, 4, 8, 2, 3, 5, 6, 7},
		{6, 7, 3, 4, 9, 5, 2, 8, 1}}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sol1b.Board[i][j] != board1Solved[i][j] {
				t.Fatalf("The solution to board %v is incorrect, r%vc%v should be %v, but got %v", board1, i+1, j+1, board1Solved[i][j], sol1b.Board[i][j])
			}
		}
	}
}
