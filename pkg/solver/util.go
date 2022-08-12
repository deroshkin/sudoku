// Package solver provides the key functionality of this sudoku solver
package solver

import (
	"golang.org/x/exp/slices"
)

// Strategy is the prototype for implementing Sudoku solving strategies.
// As the input, each strategy receives the candidates of the solver object.
// It is expected to perform any changes in place and return a single boolean
// indicating whether or not any changes have occurred.
type Strategy func([][][]uint8) bool

// Solver is the primary type for all functions.
// It consists of a collection of strategies to be used,
// the current fixed-state board (0=empty) and a list of candidates.
// The strategies will be called one at a time until one changes something,
// then the solver will start from the beginning.
type Solver struct {
	Strats []Strategy
	Board  [][]uint8
	Cands  [][][]uint8
}

// Take a board (0s used for empty cells) and a list of strategies and create
// a new solver.
func MakeSolver(board [][]uint8, strats []Strategy) *Solver {
	cands := makeCandidates(board)
	boardCopy := make([][]uint8, 9)
	for i := 0; i < 9; i++ {
		boardCopy[i] = make([]uint8, 9)
		for j := 0; j < 9; j++ {
			boardCopy[i][j] = board[i][j]
		}
	}
	stratsCopy := make([]Strategy, len(strats))
	for i := 0; i < len(strats); i++ {
		stratsCopy[i] = strats[i]
	}
	sol := Solver{
		Strats: stratsCopy,
		Board:  boardCopy,
		Cands:  cands,
	}
	return &sol
}

// Initialize the candidates for each cell, 1-9 if the cell is empty
// and just the given value if the cell is already filled.
func makeCandidates(board [][]uint8) [][][]uint8 {
	cands := make([][][]uint8, 9)
	for i := 0; i < 9; i++ {
		cands[i] = make([][]uint8, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				cands[i][j] = []uint8{board[i][j]}
			} else {
				cands[i][j] = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}
	}
	return cands
}

// Check the legality of the current board state of a solver,
// only verifies that no two copies of the same digit see each other.
func (sol *Solver) IsLegal() bool {
	for i := 0; i < 9; i++ {
		row := 0
		col := 0
		box := 0
		for j := 0; j < 9; j++ {
			if sol.Board[i][j] != 0 && row&(1<<sol.Board[i][j]) != 0 {
				return false
			} else {
				row |= 1 << sol.Board[i][j]
			}
			if sol.Board[j][i] != 0 && col&(1<<sol.Board[j][i]) != 0 {
				return false
			} else {
				col |= 1 << sol.Board[j][i]
			}
			if sol.Board[3*(i/3)+j/3][3*(i%3)+j%3] != 0 && box&(1<<sol.Board[3*(i/3)+j/3][3*(i%3)+j%3]) != 0 {
				return false
			} else {
				box |= 1 << sol.Board[3*(i/3)+j/3][3*(i%3)+j%3]
			}
		}
	}
	return true
}

// Remove the given value from the list of candidates for a given cell,
// can be used even if the value is not a candidate.
func RemoveCand(cands [][][]uint8, i, j int, val uint8) (change bool) {
	change = false
	ind := slices.Index(cands[i][j], val)
	if ind >= 0 {
		change = true
		cn := append(cands[i][j][:ind], cands[i][j][ind+1:]...)
		cands[i][j] = cn
	}
	return
}

// Check to make sure that there is at least one copy of each digit available in
// the list of candidates for each row/column/box.
func (sol *Solver) EnoughCands() bool {
	for i := 0; i < 9; i++ {
		row := 0
		col := 0
		box := 0

		for j := 0; j < 9; j++ {
			for _, v := range sol.Cands[i][j] {
				row |= 1 << (v - 1)
			}
			for _, v := range sol.Cands[j][i] {
				col |= 1 << (v - 1)
			}
			for _, v := range sol.Cands[3*(i/3)+j/3][3*(i%3)+j%3] {
				box |= 1 << (v - 1)
			}
		}
		if row < 511 || col < 511 || box < 511 {
			return false
		}
	}
	return true
}

// Check whether the sudoku is solved (assumes that the board is in a legal state)
func (sol *Solver) IsSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sol.Board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
