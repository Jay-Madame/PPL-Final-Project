package main

import "fmt"

const boardSize = 8

var fileMap = map[rune]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7}

type Square struct {
	Piece   Piece
	Marker  string
}

type Board [8][8]Square

func NewBoard() Board {
    var b Board

    // Set up pawns
    for i := 0; i < boardSize; i++ {
        b[1][i].Piece = WhitePawn{}
        b[6][i].Piece = BlackPawn{}
    }

    // Set up other pieces
    b[0][0], b[0][7] = Square{Piece: WhiteRook{}}, Square{Piece: WhiteRook{}}
    b[0][1], b[0][6] = Square{Piece: WhiteKnight{}}, Square{Piece: WhiteKnight{}}
    b[0][2], b[0][5] = Square{Piece: WhiteBishop{}}, Square{Piece: WhiteBishop{}}
    b[0][3] = Square{Piece: WhiteQueen{}}
    b[0][4] = Square{Piece: WhiteKing{}}

    b[7][0], b[7][7] = Square{Piece: BlackRook{}}, Square{Piece: BlackRook{}}
    b[7][1], b[7][6] = Square{Piece: BlackKnight{}}, Square{Piece: BlackKnight{}}
    b[7][2], b[7][5] = Square{Piece: BlackBishop{}}, Square{Piece: BlackBishop{}}
    b[7][3] = Square{Piece: BlackQueen{}}
    b[7][4] = Square{Piece: BlackKing{}}

    return b
}
func (b Board) Display() {
	fmt.Println("   a  b  c  d  e  f  g  h")
	for i := boardSize - 1; i >= 0; i-- {
		fmt.Printf("%d ", i+1)
		for j := 0; j < boardSize; j++ {
			sq := b[i][j]
			char := "_"
			if sq.Marker != "" {
				char = sq.Marker
			} else if sq.Piece != nil {
				char = sq.Piece.Symbol()
			}
			fmt.Printf("[%s]", char)
		}
		fmt.Printf(" %d\n", i+1)
	}
	fmt.Println("   a  b  c  d  e  f  g  h\n")
}
