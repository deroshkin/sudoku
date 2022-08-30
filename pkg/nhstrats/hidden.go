package nhstrats

import (
	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

// HiddenPairs is a strategy that searches for hidden pairs in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A hidden pair occurs when a pair of values is restricted to the same 2 cells in
// a row, column, or box.
func HiddenPairs(sol *solver.Solver) (changed bool) {
	return hiddenRowkTuple(sol, 2) || hiddenColkTuple(sol, 2) || hiddenBoxkTuple(sol, 2)
}

// HiddenTriples is a strategy that searches for hidden triples in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A hidden triple occurs when three values are restricted to (a subset of) the same 3 cells in
// a row, column, or box.
func HiddenTriples(sol *solver.Solver) (changed bool) {
	return hiddenRowkTuple(sol, 3) || hiddenColkTuple(sol, 3) || hiddenBoxkTuple(sol, 3)
}

// HiddenQuads is a strategy that searches for hidden quadruples in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A hidden quadruple occurs when four values are restricted to (a subset of) the same 4 cells in
// a row, column, or box.
func HiddenQuads(sol *solver.Solver) (changed bool) {
	return hiddenRowkTuple(sol, 4) || hiddenColkTuple(sol, 4) || hiddenBoxkTuple(sol, 4)
}

// hiddenRowkTuple finds naked k-tuples in rows, returns whether any changes are made
func hiddenRowkTuple(sol *solver.Solver, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		locs := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			if len(sol.Cands[i][j]) > 1 {
				for _, v := range sol.Cands[i][j] {
					locs[v] = append(locs[v], j)
				}
			}
		}
		locs = restrict(locs, k)
		matches := findMatches(locs, k, 0)
		for _, setVals := range matches {
			for j := uint8(0); j < 9; j++ {
				isPart := false
				isFull := true
				intersect := []uint8{}
				for _, v := range sol.Cands[i][j] {
					if slices.Contains(setVals, v) {
						isPart = true
						intersect = append(intersect, v)
					} else {
						isFull = false
					}
				}
				if isPart && !isFull {
					sol.Cands[i][j] = intersect
					changed = true
				}
			}
			if changed {
				return
			}
		}
	}
	return
}

// hiddenColkTuple finds naked k-tuples in columns, returns whether any changes are made
func hiddenColkTuple(sol *solver.Solver, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		locs := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			if len(sol.Cands[j][i]) > 1 {
				for _, v := range sol.Cands[j][i] {
					locs[v] = append(locs[v], j)
				}
			}
		}
		locs = restrict(locs, k)
		matches := findMatches(locs, k, 0)
		for _, setVals := range matches {
			for j := uint8(0); j < 9; j++ {
				isPart := false
				isFull := true
				intersect := []uint8{}
				for _, v := range sol.Cands[j][i] {
					if slices.Contains(setVals, v) {
						isPart = true
						intersect = append(intersect, v)
					} else {
						isFull = false
					}
				}
				if isPart && !isFull {
					sol.Cands[j][i] = intersect
					changed = true
				}
			}
			if changed {
				return
			}
		}
	}
	return
}

// hiddenBoxkTuple finds naked k-tuples in boxes, returns whether any changes are made
func hiddenBoxkTuple(sol *solver.Solver, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		locs := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			if len(sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)]) > 1 {
				for _, v := range sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					locs[v] = append(locs[v], j)
				}
			}
		}
		locs = restrict(locs, k)
		matches := findMatches(locs, k, 0)
		for _, setVals := range matches {
			for j := uint8(0); j < 9; j++ {
				isPart := false
				isFull := true
				intersect := []uint8{}
				for _, v := range sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					if slices.Contains(setVals, v) {
						isPart = true
						intersect = append(intersect, v)
					} else {
						isFull = false
					}
				}
				if isPart && !isFull {
					sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] = intersect
					changed = true
				}
			}
			if changed {
				return
			}
		}
	}
	return
}
