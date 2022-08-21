// Package interstrats provides strategies for handling candidate removal thanks to intersections,
// these are also called pointing pairs/triples and box line reductions.
package interstrats

import (
	"github.com/deroshkin/sudoku/pkg/solver"
	"golang.org/x/exp/slices"
)

// PointingSets is the strategy for handling pointing pairs/triples.
// It iterates over the boxes, and for each box checks digits 1-9 first for row pointers, then for column pointers.
// It returns true as soon as one pointing set forces a change, and false if no change is forced.
// Note: A pointing pair/triple occurs when all candidates for where to place a digit in a given box
// occur in the same row or in the same column.
// In that case, we can remove the digit from the rest of the row/column.
func PointingSets(cands [][][]uint8) (changed bool) {
	for box := uint8(0); box < 9; box++ {
		if pointers(cands, box) {
			return true
		}
	}
	return
}

func pointers(cands [][][]uint8, box uint8) (changed bool) {
	digitRows := map[uint8][]uint8{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
	digitCols := map[uint8][]uint8{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
	for i := uint8(0); i < 9; i++ {
		if len(cands[3*(box/3)+(i/3)][3*(box%3)+(i%3)]) > 1 {
			for _, v := range cands[3*(box/3)+(i/3)][3*(box%3)+(i%3)] {
				if !slices.Contains(digitRows[v], 3*(box/3)+(i/3)) {
					digitRows[v] = append(digitRows[v], 3*(box/3)+(i/3))
				}
				if !slices.Contains(digitCols[v], 3*(box%3)+(i%3)) {
					digitCols[v] = append(digitCols[v], 3*(box%3)+(i%3))
				}
			}
		}
	}
	for v := uint8(1); v < 10; v++ {
		if len(digitRows[v]) == 1 {
			i := digitRows[v][0]
			for j := uint8(0); j < 9; j++ {
				if 3*(i/3)+(j/3) != box {
					change := solver.RemoveCand(cands, int(i), int(j), v)
					changed = changed || change
				}
			}
		}
		if changed {
			return
		}
		if len(digitCols[v]) == 1 {
			i := digitCols[v][0]
			for j := uint8(0); j < 9; j++ {
				if 3*(j/3)+(i/3) != box {
					change := solver.RemoveCand(cands, int(j), int(i), v)
					changed = changed || change
				}
			}
		}
		if changed {
			return
		}
	}
	return
}
