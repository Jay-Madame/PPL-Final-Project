package main

import (
	"strings"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func parseMove(input string) (from [2]int, to [2]int, ok bool) {
	parts := strings.Split(strings.ToLower(strings.TrimSpace(input)), " to ")
	if len(parts) != 2 {
		return [2]int{}, [2]int{}, false
	}

	from = algebraicToIndex(parts[0])
	to = algebraicToIndex(parts[1])
	
	if !isValidPosition(from) || !isValidPosition(to) {
		return [2]int{}, [2]int{}, false
	}
	
	return from, to, true
}

func algebraicToIndex(pos string) [2]int {
	if len(pos) != 2 {
		return [2]int{-1, -1}
	}
	
	col, ok := fileMap[rune(pos[0])]
	if !ok {
		return [2]int{-1, -1}
	}
	
	row := int(pos[1] - '1')
	if row < 0 || row >= boardSize {
		return [2]int{-1, -1}
	}
	
	return [2]int{row, col}
}

func isValidPosition(pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < boardSize && pos[1] >= 0 && pos[1] < boardSize
}

func applyMove(b *Board, from, to [2]int, isWhite bool) bool {
	piece := b[from[0]][from[1]].Piece
	if piece == nil || piece.IsWhite() != isWhite {
		return false
	}

	valid := piece.ValidMoves(*b, from[0], from[1])
	for _, mv := range valid {
		if mv[0] == to[0] && mv[1] == to[1] {
			// Verify the move doesn't leave the king in check
			if wouldLeaveKingInCheck(b, from, to, isWhite) {
				return false
			}

			b[to[0]][to[1]].Piece = piece
			b[from[0]][from[1]].Piece = nil
			return true
		}
	}
	return false
}

func isKingInCheck(b *Board, row, col int, isWhite bool) bool {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if piece := b[r][c].Piece; piece != nil && piece.IsWhite() != isWhite {
				validMoves := piece.ValidMoves(*b, r, c)
				for _, move := range validMoves {
					if move[0] == row && move[1] == col {
						return true
					}
				}
			}
		}
	}
	return false
}

func wouldLeaveKingInCheck(b *Board, from, to [2]int, isWhite bool) bool {
	// Copy the board so we don't modify the original
	simBoard := copyBoard(b)

	// Make the move
	piece := simBoard[from[0]][from[1]].Piece
	simBoard[to[0]][to[1]].Piece = piece
	simBoard[from[0]][from[1]].Piece = nil

	// Find the king's position
	var kingRow, kingCol int
	found := false
	for row := 0; row < 8 && !found; row++ {
		for col := 0; col < 8; col++ {
			p := simBoard[row][col].Piece
			if p != nil && p.IsWhite() == isWhite {
				if (isWhite && p.Symbol() == "♔") || (!isWhite && p.Symbol() == "♚") {
					kingRow, kingCol = row, col
					found = true
					break
				}
			}
		}
	}

	// If king not found (shouldn't happen), assume unsafe
	if !found {
		return true
	}

	// Is king now in check?
	return isKingInCheck(simBoard, kingRow, kingCol, isWhite)
}

func copyBoard(b *Board) *Board {
	bCopy := NewBoard()
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if b[row][col].Piece != nil {
				bCopy[row][col].Piece = b[row][col].Piece.Clone()
			} else {
				bCopy[row][col].Piece = nil
			}
		}
	}
	return &bCopy
}

