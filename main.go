package main

import (
	"flag"
	"fmt"
)

var aiFlag = flag.Bool("pvc", false, "Computer Opponent")
var cvcFlag = flag.Bool("cvc", false, "Computer vs Computer")
var allPlayedMoves = []int{0}
var p1 player
var p2 player

func main() {
	flag.Parse()
	for _, row := range [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} {
		fmt.Println(row)
	}
	p1, p2 = generatePlayersRandomStart()
	winner := false
	for winner == false {
		winner = p1.collectPlay()
		displayGame()
		if winner == false {
			winner = p2.collectPlay()
			displayGame()
		}
	}
	if p1.win {
		fmt.Println(p1.token, " Wins.")
	} else if p2.win {
		fmt.Println(p2.token, " Wins.")
	} else {
		fmt.Println("Tie Game.")
	}
	fmt.Scanln()
}

func displayGame() {
	board := [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	for _, p := range [2]player{p1, p2} {
		for _, move := range p.moves {
			if move <= 3 {
				board[0][move-1] = p.token
			} else if move <= 6 {
				board[1][move-4] = p.token
			} else {
				board[2][move-7] = p.token
			}
		}
	}
	for _, row := range board {
		fmt.Println(row)
	}
}

func isNewMove(move int) bool {
	for _, idx := range allPlayedMoves {
		if move == idx {
			return false
		}
	}
	return true
}

func comparator(moveSet []int, checkForWin bool) (bool, int) {
	winSets := [3][3][3]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}, {{1, 5, 9}, {3, 5, 7}, {}}}
	for _, set := range winSets {
		for _, row := range set {
			matchCount := 0
			for _, val := range row {
				for _, move := range moveSet {
					if val == move {
						matchCount++
					}
				}
			}
			if checkForWin {
				if matchCount == 2 {
					for _, idx := range row {
						if isNewMove(idx) {
							return true, idx
						}
					}
				}
			}
			if matchCount == 3 {
				return true, 0
			}
		}
	}
	return false, 0
}
