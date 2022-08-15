package nhstrats

import (
	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

func NakedPairs(cands [][][]uint8) (changed bool) {
	return nakedRowkTuples(cands, 2) || nakedColkTuples(cands, 2) || nakedBoxkTuples(cands, 2)
}

func NakedTriples(cands [][][]uint8) (changed bool) {
	return nakedRowkTuples(cands, 3) || nakedColkTuples(cands, 3) || nakedBoxkTuples(cands, 3)
}

// Find naked k-tuples in rows, returns whether any changes are made
func nakedRowkTuples(cands [][][]uint8, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		vals := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			vals[j] = cands[i][j]
		}
		vals = restrict(vals, k)
		matches := findMatches(vals, k, 1)
		for _, set_cells := range matches {
			set := make([]uint8, 9)
			for _, j := range set_cells {
				for _, v := range cands[i][j] {
					if !slices.Contains(set, v) {
						set = append(set, v)
					}
				}
			}
			for j := uint8(0); j < 9; j++ {
				if !slices.Contains(set_cells, j) {
					for _, v := range set {
						if solver.RemoveCand(cands, int(i), int(j), v) {
							changed = true
						}
					}
				}
			}
			if changed {
				return
			}
		}
	}
	return
}

// Find naked k-tuples in columns, returns whether any changes are made
func nakedColkTuples(cands [][][]uint8, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		vals := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			vals[j] = cands[j][i]
		}
		vals = restrict(vals, k)
		matches := findMatches(vals, k, 1)
		for _, set_cells := range matches {
			set := make([]uint8, 9)
			for _, j := range set_cells {
				for _, v := range cands[j][i] {
					if !slices.Contains(set, v) {
						set = append(set, v)
					}
				}
			}
			for j := uint8(0); j < 9; j++ {
				if !slices.Contains(set_cells, j) {
					for _, v := range set {
						if solver.RemoveCand(cands, int(j), int(i), v) {
							changed = true
						}
					}
				}
			}
			if changed {
				return
			}
		}
	}
	return
}

// Find naked k-tuples in boxes, returns whether any changes are made
func nakedBoxkTuples(cands [][][]uint8, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		vals := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			vals[j] = cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)]
		}
		vals = restrict(vals, k)
		matches := findMatches(vals, k, 1)
		for _, set_cells := range matches {
			set := make([]uint8, 9)
			for _, j := range set_cells {
				for _, v := range cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					if !slices.Contains(set, v) {
						set = append(set, v)
					}
				}
			}
			for j := uint8(0); j < 9; j++ {
				if !slices.Contains(set_cells, j) {
					for _, v := range set {
						if solver.RemoveCand(cands, int(3*(i/3)+(j/3)), int(3*(i%3)+(j%3)), v) {
							changed = true
						}
					}
				}
			}
			if changed {
				return
			}
		}
	}
	return
}
