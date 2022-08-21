package nhstrats

import (
	"golang.org/x/exp/slices"
)

// HiddenPairs is a strategy that searches for hidden pairs in rows, columns and boxes (in that order).
// As soon as one is found, returns true. If none are found, returns false.
// Note: A hidden pair occurs when a pair of values is restricted to the same 2 cells in
// a row, column, or box.
func HiddenPairs(cands [][][]uint8) (changed bool) {
	return hiddenRowkTuple(cands, 2) || hiddenColkTuple(cands, 2) || hiddenBoxkTuple(cands, 2)
}

// hiddenRowkTuple finds naked k-tuples in rows, returns whether any changes are made
func hiddenRowkTuple(cands [][][]uint8, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		locs := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			if len(cands[i][j]) > 1 {
				for _, v := range cands[i][j] {
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
				for _, v := range cands[i][j] {
					if slices.Contains(setVals, v) {
						isPart = true
						intersect = append(intersect, v)
					} else {
						isFull = false
					}
				}
				if isPart && !isFull {
					cands[i][j] = intersect
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
func hiddenColkTuple(cands [][][]uint8, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		locs := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			if len(cands[j][i]) > 1 {
				for _, v := range cands[j][i] {
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
				for _, v := range cands[j][i] {
					if slices.Contains(setVals, v) {
						isPart = true
						intersect = append(intersect, v)
					} else {
						isFull = false
					}
				}
				if isPart && !isFull {
					cands[j][i] = intersect
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
func hiddenBoxkTuple(cands [][][]uint8, k int) (changed bool) {
	for i := uint8(0); i < 9; i++ {
		locs := map[uint8][]uint8{}
		for j := uint8(0); j < 9; j++ {
			if len(cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)]) > 1 {
				for _, v := range cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
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
				for _, v := range cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					if slices.Contains(setVals, v) {
						isPart = true
						intersect = append(intersect, v)
					} else {
						isFull = false
					}
				}
				if isPart && !isFull {
					cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] = intersect
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
