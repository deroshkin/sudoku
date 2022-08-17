package util

import (
	"fmt"

	"github.com/deroshkin/sudoku/pkg/solver"
)

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
	return true, ""
}
