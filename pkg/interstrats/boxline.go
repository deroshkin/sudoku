package interstrats

import (
	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

// BoxLineReduction is the dual strategy to pointing pairs/triples,
// it looks for rows/columns where all available locations for a digit are limited to a single box,
// and removes all other candidate locations of that digit from that box.
// If a change was made, it returns true, otherwise false.
func BoxLineReduction(sol *solver.Solver) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		if boxRowReduction(sol, i) || boxColReduction(sol, i) {
			return true
		}
	}
	return
}

// boxRowReduction checks for BoxLineReduction on rows
func boxRowReduction(sol *solver.Solver, row uint8) (changed bool) {
	digitBoxes := map[uint8][]uint8{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
	for i := uint8(0); i < 9; i++ {
		if len(sol.Cands[row][i]) > 1 {
			for _, v := range sol.Cands[row][i] {
				if !slices.Contains(digitBoxes[v], 3*(row/3)+(i/3)) {
					digitBoxes[v] = append(digitBoxes[v], 3*(row/3)+(i/3))
				}
			}
		}
	}
	for v := uint8(1); v < 10; v++ {
		if len(digitBoxes[v]) == 1 {
			box := digitBoxes[v][0]
			for i := uint8(0); i < 9; i++ {
				if 3*(box/3)+(i/3) != row {
					change := solver.RemoveCand(sol.Cands, int(3*(box/3)+(i/3)), int(3*(box%3)+(i%3)), v)
					changed = changed || change
				}
			}
		}
		if changed {
			sol.Logger.Printf("%v in row %v is restricted to box %v, removing from all other cells in the box\n", v, row, digitBoxes[v][0]+1)
			return
		}
	}
	return
}

// boxColReduction checks for BoxLineReduction on rows
func boxColReduction(sol *solver.Solver, col uint8) (changed bool) {
	digitBoxes := map[uint8][]uint8{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
	for i := uint8(0); i < 9; i++ {
		if len(sol.Cands[i][col]) > 1 {
			for _, v := range sol.Cands[i][col] {
				if !slices.Contains(digitBoxes[v], 3*(i/3)+(col/3)) {
					digitBoxes[v] = append(digitBoxes[v], 3*(i/3)+(col/3))
				}
			}
		}
	}
	for v := uint8(1); v < 10; v++ {
		if len(digitBoxes[v]) == 1 {
			box := digitBoxes[v][0]
			for i := uint8(0); i < 9; i++ {
				if 3*(box%3)+(i%3) != col {
					change := solver.RemoveCand(sol.Cands, int(3*(box/3)+(i/3)), int(3*(box%3)+(i%3)), v)
					changed = changed || change
				}
			}
		}
		if changed {
			sol.Logger.Printf("%v in column %v is restricted to box %v, removing from all other cells in the box\n", v, col, digitBoxes[v][0]+1)
			return
		}
	}
	return
}
