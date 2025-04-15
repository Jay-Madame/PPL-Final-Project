package main

import (
	"math/rand"
)

func Bot(moveChan chan string, movedChan chan struct{}, board *Board, gameOver chan bool) {
	for {
		select {
		case <-gameOver:
			return
		case <-moveChan:
			
			var moves [][2][2]int
			for row := 0; row < 8; row++ {
				for col := 0; col < 8; col++ {
					if piece := board[row][col].Piece; piece != nil && !piece.IsWhite() {
						from := [2]int{row, col}
						for _, to := range piece.ValidMoves(*board, row, col) {
							if !wouldLeaveKingInCheck(board, from, to, false) {
								moves = append(moves, [2][2]int{from, to})
							}
						}
					}
				}
			}

			if len(moves) > 0 {
				move := moves[rand.Intn(len(moves))]
				applyMove(board, move[0], move[1], false)
				
				if move[1][0] == 0 {
					board[move[1][0]][move[1][1]].Piece = BlackQueen{}
				}
			}

			movedChan <- struct{}{}
		}
	}
}