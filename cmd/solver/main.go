package main

import (
	"fmt"

	"github.com/deroshkin/sudoku/pkg/solver"
)

// For now this program requires a hard-coded board, proper io will be added later
func main() {
	board := [][]uint8{{2, 9, 0, 0, 0, 0, 7, 6, 0},
		{0, 8, 0, 0, 2, 4, 0, 9, 5},
		{7, 0, 5, 0, 0, 3, 2, 8, 0},
		{5, 0, 0, 6, 9, 0, 0, 1, 0},
		{0, 0, 0, 0, 3, 7, 6, 0, 0},
		{0, 4, 0, 0, 0, 5, 0, 3, 2},
		{0, 0, 3, 1, 0, 6, 0, 0, 9},
		{4, 6, 2, 0, 0, 0, 0, 7, 1},
		{0, 0, 9, 2, 7, 8, 0, 4, 0}}

	sol := solver.MakeSolver(board, []solver.Strategy{})
	sol.Solve()
	for _, row := range sol.Board {
		fmt.Println(row)
	}
}
