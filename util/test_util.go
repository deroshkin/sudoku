package util

import (
	"fmt"

	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

type Cell struct {
	R, C uint8
}

func SolveTester(board, answer [][]uint8, strats []solver.Strategy, solvable bool) (bool, string) {
	sol := solver.MakeSolver(board, strats)
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

func CandTester(board [][]uint8, strat solver.Strategy, expected bool,
	equal map[Cell][]uint8, contains map[Cell][]uint8, dnc map[Cell][]uint8) (bool, string) {
	sol := solver.MakeSolver(board, []solver.Strategy{})
	sol.RestrictCands()

	change := strat(sol.Cands)
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
