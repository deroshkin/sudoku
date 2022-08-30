package nhstrats

import (
	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

// NakedPairs is a strategy that searches for naked pairs in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A naked pair occurs when two cells that see each other can each have only two values,
// which allows us to remove those two values from the rest of the row/column/box.
func NakedPairs(sol *solver.Solver) (changed bool) {
	return nakedRowkTuples(sol, 2) || nakedColkTuples(sol, 2) || nakedBoxkTuples(sol, 2)
}

// NakedTriples is a strategy that searches for naked triples in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A naked triple occurs when three cells that see each other can each have (some of) only three values,
// which allows us to remove those three values from the rest of the row/column/box.
// Unlike naked pairs, the three cells need not have the same candidates (e.g. we can have cells with candidates
// 12, 13, and 23, and that would still be a valid 123 naked triple).
func NakedTriples(sol *solver.Solver) (changed bool) {
	return nakedRowkTuples(sol, 3) || nakedColkTuples(sol, 3) || nakedBoxkTuples(sol, 3)
}

// NakedQuads is a strategy that searches for naked quadruples in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A naked quadruple occurs when four cells that see each other can each have (some of) only four values,
// which allows us to remove those three values from the rest of the row/column/box.
// Similar to naked triples, the four cells need not have the same candidates.
func NakedQuads(sol *solver.Solver) (changed bool) {
	return nakedRowkTuples(sol, 4) || nakedColkTuples(sol, 4) || nakedBoxkTuples(sol, 4)
}

// nakedRowkTuples finds naked k-tuples in rows, returns whether any changes are made
func nakedRowkTuples(sol *solver.Solver, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		vals := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			vals[j] = sol.Cands[i][j]
		}
		vals = restrict(vals, k)
		matches := findMatches(vals, k, 1)
		for _, setCells := range matches {
			slices.Sort(setCells)
			set := []uint8{}
			for _, j := range setCells {
				for _, v := range sol.Cands[i][j] {
					if !slices.Contains(set, v) {
						set = append(set, v)
					}
				}
			}
			for j := uint8(0); j < 9; j++ {
				if !slices.Contains(setCells, j) {
					for _, v := range set {
						if solver.RemoveCand(sol.Cands, int(i), int(j), v) {
							changed = true
						}
					}
				}
			}
			if changed {
				slices.Sort(set)
				sol.Logger.Printf("Found a naked %v %v in row %v ( ", tuple_names[k], set, i+1)
				for _, j := range setCells {
					sol.Logger.Printf("r%vc%v ", i+1, j+1)
				}
				sol.Logger.Printf("), removing the values from other cells in the row\n")
				return
			}
		}
	}
	return
}

// nakedColkTuples finds naked k-tuples in columns, returns whether any changes are made
func nakedColkTuples(sol *solver.Solver, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		vals := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			vals[j] = sol.Cands[j][i]
		}
		vals = restrict(vals, k)
		matches := findMatches(vals, k, 1)
		for _, setCells := range matches {
			slices.Sort(setCells)
			set := []uint8{}
			for _, j := range setCells {
				for _, v := range sol.Cands[j][i] {
					if !slices.Contains(set, v) {
						set = append(set, v)
					}
				}
			}
			for j := uint8(0); j < 9; j++ {
				if !slices.Contains(setCells, j) {
					for _, v := range set {
						if solver.RemoveCand(sol.Cands, int(j), int(i), v) {
							changed = true
						}
					}
				}
			}
			if changed {
				slices.Sort(set)
				sol.Logger.Printf("Found a naked %v %v in column %v ( ", tuple_names[k], set, i+1)
				for _, j := range setCells {
					sol.Logger.Printf("r%vc%v ", j+1, i+1)
				}
				sol.Logger.Printf("), removing the values from other cells in the column\n")
				return
			}
		}
	}
	return
}

// nakedBoxkTuples finds naked k-tuples in boxes, returns whether any changes are made
func nakedBoxkTuples(sol *solver.Solver, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		vals := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			vals[j] = sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)]
		}
		vals = restrict(vals, k)
		matches := findMatches(vals, k, 1)
		for _, setCells := range matches {
			slices.Sort(setCells)
			set := []uint8{}
			for _, j := range setCells {
				for _, v := range sol.Cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					if !slices.Contains(set, v) {
						set = append(set, v)
					}
				}
			}
			for j := uint8(0); j < 9; j++ {
				if !slices.Contains(setCells, j) {
					for _, v := range set {
						if solver.RemoveCand(sol.Cands, int(3*(i/3)+(j/3)), int(3*(i%3)+(j%3)), v) {
							changed = true
						}
					}
				}
			}
			if changed {
				slices.Sort(set)
				sol.Logger.Printf("Found a naked %v %v in box %v ( ", tuple_names[k], set, i+1)
				for _, j := range setCells {
					sol.Logger.Printf("r%vc%v ", 3*(i/3)+(j/3)+1, 3*(i%3)+(j%3)+1)
				}
				sol.Logger.Printf("), removing the values from other cells in the box\n")
				return
			}
		}
	}
	return
}
