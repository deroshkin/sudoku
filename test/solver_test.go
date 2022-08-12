package sudoku_test

import (
	"testing"

	"github.com/deroshkin/sudoku/pkg/solver"
)

func TestCandGen(t *testing.T) {
	board := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 0}}

	sol := solver.MakeSolver(board, []solver.Strategy{})

	if len(sol.Cands[8][8]) != 9 {
		t.Fatalf("An empty cell should yield 9 candidates before restriction, got %v instead\n", len(sol.Cands[8][8]))
	}
}

func TestIsLegal(t *testing.T) {
	board1 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol := solver.MakeSolver(board1, []solver.Strategy{})
	if !sol.IsLegal() {
		t.Fatalf("Board %v is evaluated illegal, but is legal\n", board1)
	}

	board2 := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 7}}
	sol2 := solver.MakeSolver(board2, []solver.Strategy{})
	if !sol2.IsLegal() {
		t.Fatalf("Board %v is evaluated illegal, but is legal\n", board2)
	}

	board3 := [][]uint8{{1, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol3 := solver.MakeSolver(board3, []solver.Strategy{})
	if sol3.IsLegal() {
		t.Fatalf("Board %v is evaluated legal, but is illegal\n", board3)
	}

	board4 := [][]uint8{{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol4 := solver.MakeSolver(board4, []solver.Strategy{})
	if sol4.IsLegal() {
		t.Fatalf("Board %v is evaluated legal, but is illegal\n", board4)
	}

	board5 := [][]uint8{{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol5 := solver.MakeSolver(board5, []solver.Strategy{})
	if sol5.IsLegal() {
		t.Fatalf("Board %v is evaluated legal, but is illegal\n", board5)
	}
}

func TestRestrictCands(t *testing.T) {
	board1 := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 0}}
	sol1 := solver.MakeSolver(board1, []solver.Strategy{})
	ch1 := sol1.RestrictCands()
	if !ch1 {
		t.Fatalf("Restriction of %v should return true, but got false instead", board1)
	}
	if len(sol1.Cands[8][8]) != 1 {
		t.Fatalf("Restriction of %v should have exactly one candidate, got %v candidates: %v", board1, len(sol1.Cands[8][8]), sol1.Cands[8][8])
	}

	board2 := [][]uint8{{1, 0, 0, 5, 0, 0, 0, 0, 0},
		{4, 2, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol2 := solver.MakeSolver(board2, []solver.Strategy{})
	ch2 := sol2.RestrictCands()
	if !ch2 {
		t.Fatalf("Restriction of %v should return true, but got false instead", board2)
	}
	if len(sol2.Cands[0][2]) != 3 {
		t.Fatalf("Restriction of %v should have exactly three candidate, got %v candidates: %v", board2, len(sol2.Cands[8][8]), sol2.Cands[8][8])
	}

	board3 := [][]uint8{{1, 0, 0, 5, 0, 0, 0, 0, 0},
		{4, 2, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 0}}
	sol3 := solver.MakeSolver(board3, []solver.Strategy{})
	ch3 := sol3.RestrictCands()
	if !ch3 {
		t.Fatalf("Restriction of %v should return true, but got false instead", board3)
	}
	if len(sol3.Cands[0][2]) != 0 {
		t.Fatalf("Restriction of %v should have no candidate, got %v candidates: %v", board3, len(sol3.Cands[8][8]), sol3.Cands[8][8])
	}

	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol4 := solver.MakeSolver(board4, []solver.Strategy{})
	ch4 := sol4.RestrictCands()
	if ch4 {
		t.Fatalf("Restriction of %v should return false, but got true instead", board4)
	}
}

func TestEnoughCands(t *testing.T) {
	board1 := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 0}}
	sol1 := solver.MakeSolver(board1, []solver.Strategy{})
	sol1.RestrictCands()
	if !sol1.EnoughCands() {
		t.Fatalf("Board %v should have enough candidates, but returned not enough", board1)
	}

	board2 := [][]uint8{{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0, 0}}
	sol2 := solver.MakeSolver(board2, []solver.Strategy{})
	sol2.RestrictCands()
	if sol2.EnoughCands() {
		t.Fatalf("Board %v should not have enough candidates, but returned enough", board2)
	}

	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 3, 4, 5, 6, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol3 := solver.MakeSolver(board3, []solver.Strategy{})
	sol3.RestrictCands()
	if sol3.EnoughCands() {
		t.Fatalf("Board %v should not have enough candidates, but returned enough", board3)
	}

	board4 := [][]uint8{{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9}}
	sol4 := solver.MakeSolver(board4, []solver.Strategy{})
	sol4.RestrictCands()
	if sol4.EnoughCands() {
		t.Fatalf("Board %v should not have enough candidates, but returned enough", board4)
	}
}

func TestNakedSingles(t *testing.T) {
	board1 := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 0}}
	sol1 := solver.MakeSolver(board1, []solver.Strategy{})
	ch1a := sol1.NakedSingles()
	if ch1a {
		t.Fatalf("Naked singles on %v should return false before restriction, but got true", board1)
	}
	sol1.RestrictCands()
	ch1b := sol1.NakedSingles()
	if !ch1b {
		t.Fatalf("Naked singles on %v should return true after restriction, but got false", board1)
	}
	if sol1.Board[8][8] != 7 {
		t.Fatalf("Expected entry in r9c9 to be 7, got %v instead", sol1.Board[8][8])
	}
}

func TestIsSolved(t *testing.T) {
	board1 := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 0}}
	sol1 := solver.MakeSolver(board1, []solver.Strategy{})
	if sol1.IsSolved() {
		t.Fatalf("Board %v should return not solved, but got solved instead", board1)
	}
	sol1.RestrictCands()
	if sol1.IsSolved() {
		t.Fatalf("Board %v should return not solved after restriction, but got solved instead", board1)
	}
	sol1.NakedSingles()
	if !sol1.IsSolved() {
		t.Fatalf("Board %v should return solved after restriction and naked singles, but got not solved instead", board1)
	}
}

func TestSolve(t *testing.T) {
	board1 := [][]uint8{{6, 7, 2, 1, 4, 5, 3, 9, 8},
		{1, 4, 5, 9, 8, 3, 6, 7, 2},
		{3, 8, 9, 7, 6, 2, 4, 5, 1},
		{2, 6, 3, 5, 7, 4, 8, 1, 9},
		{9, 5, 8, 6, 2, 1, 7, 4, 3},
		{7, 1, 4, 3, 9, 8, 5, 2, 6},
		{5, 9, 7, 2, 3, 6, 1, 8, 4},
		{4, 2, 6, 8, 1, 7, 9, 3, 5},
		{8, 3, 1, 4, 5, 9, 2, 6, 0}}
	sol1 := solver.MakeSolver(board1, []solver.Strategy{})
	if !sol1.Solve() {
		t.Fatalf("Board %v should be solvable with no strategies, but got not solvable", board1)
	}
	if sol1.Board[8][8] != 7 {
		t.Fatalf("The solution to board %v should have 7 in r9c9, but got %v instead", board1, sol1.Board[8][8])
	}

	board2 := [][]uint8{{2, 9, 0, 0, 0, 0, 7, 6, 0},
		{0, 8, 0, 0, 2, 4, 0, 9, 5},
		{7, 0, 5, 0, 0, 3, 2, 8, 0},
		{5, 0, 0, 6, 9, 0, 0, 1, 0},
		{0, 0, 0, 0, 3, 7, 6, 0, 0},
		{0, 4, 0, 0, 0, 5, 0, 3, 2},
		{0, 0, 3, 1, 0, 6, 0, 0, 9},
		{4, 6, 2, 0, 0, 0, 0, 7, 1},
		{0, 0, 9, 2, 7, 8, 0, 4, 0}}
	board2Solved := [][]uint8{{2, 9, 4, 5, 8, 1, 7, 6, 3},
		{3, 8, 6, 7, 2, 4, 1, 9, 5},
		{7, 1, 5, 9, 6, 3, 2, 8, 4},
		{5, 3, 8, 6, 9, 2, 4, 1, 7},
		{9, 2, 1, 4, 3, 7, 6, 5, 8},
		{6, 4, 7, 8, 1, 5, 9, 3, 2},
		{8, 7, 3, 1, 4, 6, 5, 2, 9},
		{4, 6, 2, 3, 5, 9, 8, 7, 1},
		{1, 5, 9, 2, 7, 8, 3, 4, 6}}
	sol2 := solver.MakeSolver(board2, []solver.Strategy{})
	if !sol2.Solve() {
		t.Fatalf("Board %v should be solvable with no strategies, but got not solvable", board2)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sol2.Board[i][j] != board2Solved[i][j] {
				t.Fatalf("The solution to board %v should have %v in r%vc%v, but got %v instead", board2, board2Solved[i][j], i+1, j+1, sol2.Board[i][j])
			}
		}
	}

	board3 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol3 := solver.MakeSolver(board3, []solver.Strategy{})
	if sol3.Solve() {
		t.Fatalf("Board %v should be not solvable, but got solvable", board3)
	}

	board4 := [][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 3, 4, 5, 6, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol4 := solver.MakeSolver(board4, []solver.Strategy{})
	if sol4.Solve() {
		t.Fatalf("Board %v should be not solvable, but got solvable", board4)
	}

	board5 := [][]uint8{{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	sol5 := solver.MakeSolver(board5, []solver.Strategy{})
	if sol5.Solve() {
		t.Fatalf("Board %v should be not solvable, but got solvable", board5)
	}
}
