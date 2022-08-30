// Package util provides supporting functions for other packages and tests
// that are not meant to be invoked directly otherwise.
package util

import (
	"fmt"
	"io"
	"log"

	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

// Cell is a simple tuple to encode cell coordinates.
type Cell struct {
	R, C uint8
}

// SolveTester is a utility function for tests that checks whether a given sudoku board is solvable
// using the provided strategies, and if it is solvable, that the answer matches the provided one.
// As input it takes the initial board, the answer (can be [][]uint8{} if the board should not be solvable),
// a list of strategies to be used and a boolean to idicate whether the board should be solvable.
// The return is either (true, "") if the test passed, or (false, error_message) if the test failed.
// The expected use is:
// res, msg = SolveTester(...)
//
//	if !res {
//	    t.Fatalf(msg)
//	}
func SolveTester(board, answer [][]uint8, strats []solver.Strategy, solvable bool) (bool, string) {
	sol := solver.MakeSolver(board, strats, log.New(io.Discard, "", 0))
	solved := sol.Solve()
	if !solvable && solved {
		return false, fmt.Sprintf("Board %v should not be solvable using the provided strategies, but was solved", board)
	} else if solvable && !solved {
		return false, fmt.Sprintf("Board %v should be solvable using the provided strategies, but was not solved", board)
	} else if solvable {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if sol.Board[i][j] != answer[i][j] {
					return false, fmt.Sprintf("The solution of board %v does not match the expected answer, got %v in r%vc%v, expected %v",
						board, sol.Board[i][j], i+1, j+1, answer[i][j])
				}
			}
		}
	}
	return true, "Test successful"
}

// CandTester is a utility function for tests to check whether the given strategy works correctly.
// As input, this function takes a starting board, a strategy, expected outcome of the strategy (change/no change),
// and three maps Cell -> []uint8:
// equal for when the values in the indicated cell must equal the provided slice
// contains for when the values in the indicated cell must contain the provided values
// dnc for when the values in the indicated cell must not contain the provided values.
// The return is either (true, "") if the test passed, or (false, error_message) if the test failed.
// The expected use is:
// res, msg = CandTester(...)
//
//	if !res {
//	    t.Fatalf(msg)
//	}
func CandTester(board [][]uint8, strat solver.Strategy, expected bool,
	equal map[Cell][]uint8, contains map[Cell][]uint8, dnc map[Cell][]uint8) (bool, string) {
	sol := solver.MakeSolver(board, []solver.Strategy{}, log.New(io.Discard, "", 0))
	sol.RestrictCands()

	change := strat(sol)
	if !expected && change {
		return false, fmt.Sprintf("Strategy %v made changes when none were expected", strat)
	} else if expected && !change {
		return false, fmt.Sprintf("Strategy %v made no changes when some were expected", strat)
	} else if expected {
		for cell, vals := range equal {
			if !slices.Equal(sol.Cands[cell.R][cell.C], vals) {
				return false, fmt.Sprintf("Expected r%vc%v to have candidates %v, but got %v instead",
					cell.R+1, cell.C+1, vals, sol.Cands[cell.R][cell.C])
			}
		}
		for cell, vals := range contains {
			for _, v := range vals {
				if !slices.Contains(sol.Cands[cell.R][cell.C], v) {
					return false, fmt.Sprintf("Expected r%vc%v to contain candidates %v, but got %v instead",
						cell.R+1, cell.C+1, vals, sol.Cands[cell.R][cell.C])
				}
			}
		}
		for cell, vals := range dnc {
			for _, v := range vals {
				if slices.Contains(sol.Cands[cell.R][cell.C], v) {
					return false, fmt.Sprintf("Expected candidates %v to be missing from r%vc%v, but have candidates %v",
						vals, cell.R+1, cell.C+1, sol.Cands[cell.R][cell.C])
				}
			}
		}
	}
	return true, "Test successful"
}
