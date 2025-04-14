package main

type Piece interface {
    Symbol() string
    IsWhite() bool
    ValidMoves(b Board, row, col int) [][2]int
	Clone() Piece
}

// White pieces
type WhitePawn struct{}
func (p WhitePawn) Symbol() string { return "♙" }
func (p WhitePawn) IsWhite() bool { return true }
func (p WhitePawn) ValidMoves(b Board, row, col int) [][2]int {
    var moves [][2]int
    // Forward move
    if row+1 < 8 && b[row+1][col].Piece == nil {
        moves = append(moves, [2]int{row+1, col})
        // Double move from starting position
        if row == 1 && b[row+2][col].Piece == nil {
            moves = append(moves, [2]int{row+2, col})
        }
    }
    // Captures
    for _, dc := range []int{-1, 1} {
        if col+dc >= 0 && col+dc < 8 && row+1 < 8 {
            if target := b[row+1][col+dc].Piece; target != nil && !target.IsWhite() {
                moves = append(moves, [2]int{row+1, col+dc})
            }
        }
    }
    return moves
}
func (w WhitePawn) Clone() Piece { return WhitePawn{}}

type WhiteKing struct{}
func (k WhiteKing) Symbol() string { return "♔" }
func (k WhiteKing) IsWhite() bool { return true }
func (k WhiteKing) ValidMoves(b Board, row, col int) [][2]int {
    var moves [][2]int
    for dr := -1; dr <= 1; dr++ {
        for dc := -1; dc <= 1; dc++ {
            if dr == 0 && dc == 0 { continue }
            newRow, newCol := row+dr, col+dc
            if newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8 {
                if b[newRow][newCol].Piece == nil || !b[newRow][newCol].Piece.IsWhite() {
                    moves = append(moves, [2]int{newRow, newCol})
                }
            }
        }
    }
    return moves
}
func (w WhiteKing) Clone() Piece { return WhiteKing{}}

// Black pieces
type BlackPawn struct{}
func (p BlackPawn) Symbol() string { return "♟︎" }
func (p BlackPawn) IsWhite() bool { return false }
func (p BlackPawn) ValidMoves(b Board, row, col int) [][2]int {
    var moves [][2]int
    // Forward move
    if row-1 >= 0 && b[row-1][col].Piece == nil {
        moves = append(moves, [2]int{row-1, col})
        // Double move from starting position
        if row == 6 && b[row-2][col].Piece == nil {
            moves = append(moves, [2]int{row-2, col})
        }
    }
    // Captures
    for _, dc := range []int{-1, 1} {
        if col+dc >= 0 && col+dc < 8 && row-1 >= 0 {
            if target := b[row-1][col+dc].Piece; target != nil && target.IsWhite() {
                moves = append(moves, [2]int{row-1, col+dc})
            }
        }
    }
    return moves
}
func (w BlackPawn) Clone() Piece { return BlackPawn{}}

type BlackKing struct{}
func (k BlackKing) Symbol() string { return "♚" }
func (k BlackKing) IsWhite() bool { return false }
func (k BlackKing) ValidMoves(b Board, row, col int) [][2]int {
    var moves [][2]int
    for dr := -1; dr <= 1; dr++ {
        for dc := -1; dc <= 1; dc++ {
            if dr == 0 && dc == 0 { continue }
            newRow, newCol := row+dr, col+dc
            if newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8 {
                if b[newRow][newCol].Piece == nil || b[newRow][newCol].Piece.IsWhite() {
                    moves = append(moves, [2]int{newRow, newCol})
                }
            }
        }
    }
    return moves
}
func (w BlackKing) Clone() Piece { return BlackKing{}}

// White Rook
type WhiteRook struct{}
func (r WhiteRook) Symbol() string { return "♖" }
func (r WhiteRook) IsWhite() bool { return true }
func (r WhiteRook) ValidMoves(b Board, row, col int) [][2]int {
	var moves [][2]int
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		for i := 1; i < 8; i++ {
			nr, nc := row+i*d[0], col+i*d[1]
			if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
				break
			}
			target := b[nr][nc].Piece
			if target == nil {
				moves = append(moves, [2]int{nr, nc})
			} else {
				if !target.IsWhite() {
					moves = append(moves, [2]int{nr, nc})
				}
				break
			}
		}
	}
	return moves
}
func (w WhiteRook) Clone() Piece { return WhiteRook{}}

// Black Rook
type BlackRook struct{}
func (r BlackRook) Symbol() string { return "♜" }
func (r BlackRook) IsWhite() bool { return false }
func (r BlackRook) ValidMoves(b Board, row, col int) [][2]int {
	var moves [][2]int
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		for i := 1; i < 8; i++ {
			nr, nc := row+i*d[0], col+i*d[1]
			if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
				break
			}
			target := b[nr][nc].Piece
			if target == nil {
				moves = append(moves, [2]int{nr, nc})
			} else {
				if target.IsWhite() {
					moves = append(moves, [2]int{nr, nc})
				}
				break
			}
		}
	}
	return moves
}
func (w BlackRook) Clone() Piece { return BlackRook{}}

// White Knight
type WhiteKnight struct{}
func (n WhiteKnight) Symbol() string { return "♘" }
func (n WhiteKnight) IsWhite() bool { return true }
func (n WhiteKnight) ValidMoves(b Board, row, col int) [][2]int {
	deltas := [][2]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	var moves [][2]int
	for _, d := range deltas {
		nr, nc := row+d[0], col+d[1]
		if nr >= 0 && nr < 8 && nc >= 0 && nc < 8 {
			target := b[nr][nc].Piece
			if target == nil || !target.IsWhite() {
				moves = append(moves, [2]int{nr, nc})
			}
		}
	}
	return moves
}
func (w WhiteKnight) Clone() Piece { return WhiteKnight{}}

// Black Knight
type BlackKnight struct{}
func (n BlackKnight) Symbol() string { return "♞" }
func (n BlackKnight) IsWhite() bool { return false }
func (n BlackKnight) ValidMoves(b Board, row, col int) [][2]int {
	deltas := [][2]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	var moves [][2]int
	for _, d := range deltas {
		nr, nc := row+d[0], col+d[1]
		if nr >= 0 && nr < 8 && nc >= 0 && nc < 8 {
			target := b[nr][nc].Piece
			if target == nil || target.IsWhite() {
				moves = append(moves, [2]int{nr, nc})
			}
		}
	}
	return moves
}
func (w BlackKnight) Clone() Piece { return BlackKnight{}}

// White Bishop
type WhiteBishop struct{}
func (b WhiteBishop) Symbol() string { return "♗" }
func (b WhiteBishop) IsWhite() bool { return true }
func (b WhiteBishop) ValidMoves(board Board, row, col int) [][2]int {
	return diagonalMoves(board, row, col, true)
}
func (w WhiteBishop) Clone() Piece { return WhiteBishop{}}

// Black Bishop
type BlackBishop struct{}
func (b BlackBishop) Symbol() string { return "♝" }
func (b BlackBishop) IsWhite() bool { return false }
func (b BlackBishop) ValidMoves(board Board, row, col int) [][2]int {
	return diagonalMoves(board, row, col, false)
}
func (w BlackBishop) Clone() Piece { return BlackBishop{}}

// Helper
func diagonalMoves(b Board, row, col int, white bool) [][2]int {
	var moves [][2]int
	directions := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, d := range directions {
		for i := 1; i < 8; i++ {
			nr, nc := row+i*d[0], col+i*d[1]
			if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
				break
			}
			target := b[nr][nc].Piece
			if target == nil {
				moves = append(moves, [2]int{nr, nc})
			} else {
				if target.IsWhite() != white {
					moves = append(moves, [2]int{nr, nc})
				}
				break
			}
		}
	}
	return moves
}

// White Queen
type WhiteQueen struct{}
func (q WhiteQueen) Symbol() string { return "♕" }
func (q WhiteQueen) IsWhite() bool { return true }
func (q WhiteQueen) ValidMoves(b Board, row, col int) [][2]int {
	return append(
		diagonalMoves(b, row, col, true),
		straightMoves(b, row, col, true)...,
	)
}
func (w WhiteQueen) Clone() Piece { return WhiteQueen{}}

// Black Queen
type BlackQueen struct{}
func (q BlackQueen) Symbol() string { return "♛" }
func (q BlackQueen) IsWhite() bool { return false }
func (q BlackQueen) ValidMoves(b Board, row, col int) [][2]int {
	return append(
		diagonalMoves(b, row, col, false),
		straightMoves(b, row, col, false)...,
	)
}
func (w BlackQueen) Clone() Piece { return BlackQueen{}}

// Helper
func straightMoves(b Board, row, col int, white bool) [][2]int {
	var moves [][2]int
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		for i := 1; i < 8; i++ {
			nr, nc := row+i*d[0], col+i*d[1]
			if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
				break
			}
			target := b[nr][nc].Piece
			if target == nil {
				moves = append(moves, [2]int{nr, nc})
			} else {
				if target.IsWhite() != white {
					moves = append(moves, [2]int{nr, nc})
				}
				break
			}
		}
	}
	return moves
}
