// The nhstrats package provides strategies for naked/hidden doubles-quads and
// hidden singles (naked singles are handles by the solver package)
package nhstrats

import "golang.org/x/exp/slices"

// Restrict a map {j: cands @ j} or {v: locs for v} to only those elements
// that appear more than once and at most k<=4 times
func restrict(vals map[uint8][]uint8, k int) (res map[uint8][]uint8) {
	res = make(map[uint8][]uint8)
	for key, val := range vals {
		if len(val) > 1 && len(val) <= k {
			res[key] = val
		}
	}
	return
}

// In an already restricted map, find all sets of k<=4 matching entries and
// return them, or return {} if no match found
func findMatches(vals map[uint8][]uint8, k int, min int) (res [][]uint8) {
	inverted := make(map[uint16][]uint8, 9)
	for key, val := range vals {
		tot := uint16(0)
		for _, v := range val {
			tot += 1 << v
		}
		if len(val) == k {
			inverted[tot] = append(inverted[tot], key)
			if len(inverted[tot]) == k {
				res = append(res, inverted[tot])
			}
		} else if len(val) == k-1 {
			for i := min; i < min+9; i++ {
				if !slices.Contains(val, uint8(i)) {
					inverted[tot+1<<i] = append(inverted[tot+1<<i], key)
					if len(inverted[tot+1<<i]) == k {
						res = append(res, inverted[tot+1<<i])
					}
				}
			}
		} else { // len(val) == k-2
			for i := min; i < min+8; i++ {
				if !slices.Contains(val, uint8(i)) {
					for j := i + 1; j < min+9; j++ {
						if !slices.Contains(val, uint8(j)) {
							inverted[tot+1<<i+1<<j] = append(inverted[tot+1<<i+1<<j], key)
							if len(inverted[tot+1<<i+1<<j]) == k {
								res = append(res, inverted[tot+1<<i+1<<j])
							}
						}
					}
				}
			}
		}
	}
	return
}
