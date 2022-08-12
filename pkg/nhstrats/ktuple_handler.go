// The nhstrats package provides strategies for naked/hidden doubles-quads and
// hidden singles (naked singles are handles by the solver package)
package nhstrats

// Restrict a map {j: cands @ j} or {v: locs for v} to only those elements that appear exactly k times
func restrict(vals map[uint8][]uint8, k int) (res map[uint8][]uint8) {
	for key, val := range vals {
		if len(val) == k {
			res[key] = val
		}
	}
	return
}

// In an already restricted map, find k matching ones and return them, or return {} if no match found
func findMatch(vals map[uint8][]uint8, k int) []uint8 {
	inverted := map[uint16][]uint8{}
	for key, val := range vals {
		tot := uint16(0)
		for _, v := range val {
			tot += 1 << v
		}
		inverted[tot] = append(inverted[tot], key)
		if len(inverted[tot]) == k {
			return inverted[tot]
		}
	}
	return []uint8{}
}
