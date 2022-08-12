package solver

// Remove candidates that see a copy of the digit in the same row/column/box,
// returns whether anything has been changed.
func (sol *Solver) RestrictCands() (change bool) {
	change = false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sol.Board[i][j] == 0 {
				for k := 0; k < 9; k++ {
					if sol.Board[i][k] != 0 {
						ch := RemoveCand(sol.Cands, i, j, sol.Board[i][k])
						change = change || ch
					}
					if sol.Board[k][j] != 0 {
						ch := RemoveCand(sol.Cands, i, j, sol.Board[k][j])
						change = change || ch
					}
					box := 3*(i/3) + j/3
					if sol.Board[3*(box/3)+k/3][3*(box%3)+k%3] != 0 {
						ch := RemoveCand(sol.Cands, i, j, sol.Board[3*(box/3)+k/3][3*(box%3)+k%3])
						change = change || ch
					}
				}
			}
		}
	}
	return
}

// Search for naked singles and move any that are found to the board,
// returns whether anything has been changed.
// Note: A naked single is the term for sudoku cells where there is exactly one
// candidate value remaining.
func (sol *Solver) NakedSingles() (change bool) {
	change = false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sol.Board[i][j] == 0 && len(sol.Cands[i][j]) == 1 {
				sol.Board[i][j] = sol.Cands[i][j][0]
				change = true
			}
		}
	}
	return
}

// The main method of the solver package. This method iterates over all
// the provided strategies and returns true if the sudoku is uniquely solvable
// using those strategies, false otherwise. The solution can be accessed via
// sol.Board.
func (sol *Solver) Solve() bool {
	if !sol.IsLegal() || !sol.EnoughCands() {
		return false
	}
	for true {
		ch1 := sol.RestrictCands()
		ch2 := sol.NakedSingles()
		if ch1 || ch2 {
			if !sol.IsLegal() || !sol.EnoughCands() {
				return false
			} else if sol.IsSolved() {
				return true
			} else {
				continue
			}
		}
		for _, strat := range sol.Strats {
			if strat(sol.Cands) {
				if !sol.IsLegal() || !sol.EnoughCands() {
					return false
				} else if sol.IsSolved() {
					return true
				} else {
					continue
				}
			}
		}
		break
	}
	return false
}
