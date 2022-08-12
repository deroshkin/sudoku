// The nhstrats package provides strategies for naked/hidden doubles-quads and
// hidden singles (naked singles are handles by the solver package)
package nhstrats

// Find the hidden singles in the provided sudoku candidates.
// The first time one is found, change the candidates and return true,
// if none are found, return false.
// Note: hidden singles are the digits that can only appear in one spot in a row/column/box
func HiddenSingles(cands [][][]uint8) (changed bool) {
	return hiddenRowSingles(cands) || hiddenColSingles(cands) || hiddenBoxSingles(cands)
}

// Search for hidden singles in each row
func hiddenRowSingles(cands [][][]uint8) (changed bool) {
	changed = false
	for i := 0; i < 9; i++ {
		vals := map[uint8][]int{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
		for j := 0; j < 9; j++ {
			if len(cands[i][j]) > 1 {
				for _, v := range cands[i][j] {
					vals[v] = append(vals[v], j)
				}
			}
		}
		for v := uint8(1); v <= 9; v++ {
			if len(vals[v]) == 1 {
				cands[i][vals[v][0]] = []uint8{v}
				return true
			}
		}
	}
	return
}

// Search fo hidden singles in each column
func hiddenColSingles(cands [][][]uint8) (changed bool) {
	changed = false
	for i := 0; i < 9; i++ {
		// Column i
		vals := map[uint8][]int{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
		for j := 0; j < 9; j++ {
			if len(cands[j][i]) > 1 {
				for _, v := range cands[j][i] {
					vals[v] = append(vals[v], j)
				}
			}
		}
		for v := uint8(1); v <= 9; v++ {
			if len(vals[v]) == 1 {
				cands[vals[v][0]][i] = []uint8{v}
				return true
			}
		}
	}
	return
}

// Search for hidden singles in each box
func hiddenBoxSingles(cands [][][]uint8) (changed bool) {
	changed = false
	for i := 0; i < 9; i++ {
		// Box i
		vals := map[uint8][]int{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
		for j := 0; j < 9; j++ {
			if len(cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)]) > 1 {
				for _, v := range cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] {
					vals[v] = append(vals[v], j)
				}
			}
		}
		for v := uint8(1); v <= 9; v++ {
			if len(vals[v]) == 1 {
				j := vals[v][0]
				cands[3*(i/3)+(j/3)][3*(i%3)+(j%3)] = []uint8{v}
				return true
			}
		}
	}
	return
}
