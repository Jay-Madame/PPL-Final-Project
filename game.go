package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartGame() {
	board := NewBoard()
	playerTurn := true
	moveChan := make(chan string, 1)
	movedChan := make(chan struct{}, 1)
	gameOver := make(chan bool, 1)

	go Bot(moveChan, movedChan, &board, gameOver)

	for {
		board.Display()
		
		if isCheckmate(&board, !playerTurn) {
			if playerTurn {
				fmt.Println("Checkmate! You win!")
			} else {
				fmt.Println("Checkmate! Bot wins!")
			}
			gameOver <- true
			break
		} else if isStalemate(&board, !playerTurn) {
			fmt.Println("Stalemate! Game drawn.")
			gameOver <- true
			break
		}

		if playerTurn {
			fmt.Print("Enter move (e.g. e2 to e4): ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "quit" {
				fmt.Println("Game ended by player.")
				gameOver <- true
				break
			}

			from, to, ok := parseMove(input)
			if !ok {
				fmt.Println("Invalid input. Please use format like 'e2 to e4'")
				continue
			}

			if !isValidMove(&board, from, to, true) {
				fmt.Println("Illegal move.")
				continue
			}

			if !applyMove(&board, from, to, true) {
				fmt.Println("Move failed.")
				continue
			}

			if to[0] == 7 {
				board[to[0]][to[1]].Piece = WhiteQueen{}
			}

			playerTurn = false
			moveChan <- "bot"
			<-movedChan
		} else {
			playerTurn = true
		}
	}
}

func isValidMove(b *Board, from, to [2]int, isWhite bool) bool {
	piece := b[from[0]][from[1]].Piece
	if piece == nil || piece.IsWhite() != isWhite {
		return false
	}

	validMoves := piece.ValidMoves(*b, from[0], from[1])
	moveFound := false
	for _, mv := range validMoves {
		if mv[0] == to[0] && mv[1] == to[1] {
			moveFound = true
			break
		}
	}
	if !moveFound {
		return false
	}

	return !wouldLeaveKingInCheck(b, from, to, isWhite)
}

func isCheckmate(b *Board, isWhite bool) bool {
	var kingRow, kingCol int
	found := false
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if piece := b[row][col].Piece; piece != nil && piece.IsWhite() == isWhite {
				if (isWhite && piece.Symbol() == "♔") || (!isWhite && piece.Symbol() == "♚") {
					kingRow, kingCol = row, col
					found = true
					break
				}
			}
		}
		if found {
			break
		}
	}

	if !found {
		return false
	}

	if !isKingInCheck(b, kingRow, kingCol, isWhite) {
		return false
	}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if piece := b[row][col].Piece; piece != nil && piece.IsWhite() == isWhite {
				from := [2]int{row, col}
				for _, to := range piece.ValidMoves(*b, row, col) {
					if !wouldLeaveKingInCheck(b, from, to, isWhite) {
						return false
					}
				}
			}
		}
	}

	return true
}

func isStalemate(b *Board, isWhite bool) bool {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if piece := b[row][col].Piece; piece != nil && piece.IsWhite() == isWhite {
				from := [2]int{row, col}
				for _, to := range piece.ValidMoves(*b, row, col) {
					if !wouldLeaveKingInCheck(b, from, to, isWhite) {
						return false
					}
				}
			}
		}
	}

	return true
}