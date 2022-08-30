// Package nhstrats provides strategies for naked/hidden doubles-quads and
// hidden singles (naked singles are handles by the solver package)
package nhstrats

import "github.com/deroshkin/sudoku/pkg/solver"

// HiddenSingles finds the hidden singles in the provided sudoku candidates.
// The first time one is found, change the candidates and return true,
// if none are found, return false. Order of priority is rows, then columns, then boxes.
// Note: hidden singles are the digits that can only appear in one spot in a row/column/box.
func HiddenSingles(sol *solver.Solver) (changed bool) {
	return hiddenRowSingles(sol) || hiddenColSingles(sol) || hiddenBoxSingles(sol)
}

// hiddenRowSingles searches for hidden singles in each row
func hiddenRowSingles(sol *solver.Solver) (changed bool) {
	changed = false
	for i := 0; i < 9; i++ {
		vals := map[uint8][]int{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
		for j := 0; j < 9; j++ {
			if len(sol.Cands[i][j]) > 1 {
				for _, v := range sol.Cands[i][j] {
					vals[v] = append(vals[v], j)
				}
			}
		}
		for v := uint8(1); v <= 9; v++ {
			if len(vals[v]) == 1 {
				sol.Cands[i][vals[v][0]] = []uint8{v}
				sol.Logger.Printf("Found a hidden single %v in row %v: r%vc%v, removing all other candidates from the cell\n", v, i+1, i+1, vals[v][0]+1)
				changed = true
			}
		}
	}
	return
}

// hiddenColSingles searches for hidden singles in each column
func hiddenColSingles(sol *solver.Solver) (changed bool) {
	changed = false
	for i := 0; i < 9; i++ {
		// Column i
		vals := map[uint8][]int{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
		for j := 0; j < 9; j++ {
			if len(sol.Cands[j][i]) > 1 {
				for _, v := range sol.Cands[j][i] {
					vals[v] = append(vals[v], j)
				}
			}
		}
		for v := uint8(1); v <= 9; v++ {
			if len(vals[v]) == 1 {
				sol.Cands[vals[v][0]][i] = []uint8{v}
				sol.Logger.Printf("Found a hidden single %v in column %v: r%vc%v, removing all other candidates from the cell\n", v, i+1, vals[v][0]+1, i+1)
				changed = true
			}
		}
	}
	return
}

// hiddenBoxSingles searches for hidden singles in each box
func hiddenBoxSingles(sol *solver.Solver) (changed bool) {
	changed = false
	for i := 0; i < 9; i++ {
		// Box i
		vals := map[uint8][]int{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
		for j := 0; j < 9; j++ {
			if len(sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)]) > 1 {
				for _, v := range sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					vals[v] = append(vals[v], j)
				}
			}
		}
		for v := uint8(1); v <= 9; v++ {
			if len(vals[v]) == 1 {
				j := vals[v][0]
				sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] = []uint8{v}
				sol.Logger.Printf("Found a hidden single %v in box %v: r%vc%v, removing all other candidates from the cell\n", v, i+1, 3*(i/3)+(j/3)+1, 3*(i%3)+(j%3)+1)
				changed = true
			}
		}
	}
	return
}
