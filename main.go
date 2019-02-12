package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type player struct {	// NOTE: REVISED AND MOVED TO REV/
// 	moveset []int
// 	turnCnt int
// 	win     bool
// 	char    string
// 	ai      bool
// }

// var aiFlag = flag.Bool("ai", false, "Computer opponent")

// func main() {
// 	flag.Parse()
// 	allMoves := []int{0}
// 	indexMap := [3][3]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}
// 	for _, row := range indexMap {
// 		fmt.Println(row)
// 	}
// 	p1, p2 := generatePlayersRandomStart(aiFlag)
// 	winFlag := false
// 	for winFlag == false { // Turn management
// 		winFlag = p1.collectPlayAndCheckWin(&allMoves, p2.moveset)
// 		displayGame(p1, p2)
// 		if winFlag == false {
// 			winFlag = p2.collectPlayAndCheckWin(&allMoves, p1.moveset)
// 			displayGame(p1, p2)
// 		}
// 	}
// 	if p1.win {
// 		fmt.Println(p1.char, " Wins!")
// 	} else if p2.win {
// 		fmt.Println(p2.char, " Wins!")
// 	} else {
// 		fmt.Println("Tie Game!")
// 	}
// 	// For compiled code, to avoid terminal auto-close
// 	var null string
// 	fmt.Scanln(&null)
// }

// func generatePlayersRandomStart(aiFlag *bool) (player, player) {
// 	rand.Seed(time.Now().UnixNano())
// 	randRes := rand.Intn(2)
// 	p1Char, p2Char := "X", "O"
// 	if randRes == 1 {
// 		p1Char, p2Char = "O", "X"
// 	}
// 	p1 := player{[]int{}, 0, false, p1Char, false}
// 	p2 := player{[]int{}, 0, false, p2Char, false}
// 	if *aiFlag {
// 		anoRes := rand.Intn(2)
// 		if anoRes == 1 {
// 			p1.ai = true
// 		} else {
// 			p2.ai = true
// 		}
// 	}
// 	return p1, p2
// }

// func displayGame(p1 player, p2 player) {
// 	charset := [3][3]string{
// 		{"-", "-", "-"},
// 		{"-", "-", "-"},
// 		{"-", "-", "-"},
// 	}
// 	displayHelper(&charset, p1.moveset, p1.char)
// 	displayHelper(&charset, p2.moveset, p2.char)
// 	for _, row := range charset {
// 		fmt.Println(row)
// 	}
// }

// func displayHelper(charset *[3][3]string, moveset []int, playChar string) {
// 	// Place string characters ('X', 'O') on each players move index.
// 	for _, move := range moveset {
// 		if move <= 3 {
// 			charset[0][move-1] = playChar
// 		} else if move <= 6 {
// 			charset[1][move-4] = playChar
// 		} else {
// 			charset[2][move-7] = playChar
// 		}
// 	}
// }

// func (p *player) collectPlayAndCheckWin(allMoves *[]int, oppMoves []int) bool {
// 	xWinset := [3][3]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}
// 	yWinset := [3][3]int{
// 		{1, 4, 7},
// 		{2, 5, 8},
// 		{3, 6, 9},
// 	}
// 	dWinset := [3][3]int{
// 		{1, 5, 9},
// 		{3, 5, 7},
// 		{},
// 	}
// 	moveIndex := 0
// 	if p.ai {
// 		moveIndex = p.magic(oppMoves, allMoves)
// 		fmt.Println("Computer moves: " + strconv.Itoa(moveIndex))
// 	} else { // Human input
// 		passCnt := 0
// 		moveIndex = inputHelper(p.char)
// 		for passCnt != len(*allMoves) { // Loop until valid move played
// 			for _, move := range *allMoves {
// 				if move == moveIndex {
// 					fmt.Println("Move already played.")
// 					passCnt = 0
// 					moveIndex = inputHelper(p.char)
// 					break
// 				} else {
// 					// Limit move input 0-9
// 					if moveIndex <= 9 {
// 						passCnt++
// 					} else {
// 						moveIndex = 0
// 					}
// 				}
// 			}
// 		}
// 	}
// 	*allMoves = append(*allMoves, moveIndex)
// 	fmt.Println(*allMoves) // NOTE: TEMP
// 	p.moveset = append(p.moveset, moveIndex)
// 	p.turnCnt++
// 	if p.turnCnt > 2 {
// 		dWin, _ := comparator(dWinset, p.moveset, false, allMoves)
// 		xWin, _ := comparator(xWinset, p.moveset, false, allMoves)
// 		yWin, _ := comparator(yWinset, p.moveset, false, allMoves)
// 		if dWin || xWin || yWin {
// 			p.win = true
// 			return true
// 		} else if len(*allMoves) == 10 { // Tie game
// 			return true
// 		}
// 	}
// 	return false
// }

func (p *player) magic(oppMoves []int, allMoves *[]int) int {
	xWinset := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	yWinset := [3][3]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	dWinset := [3][3]int{
		{1, 5, 9},
		{3, 5, 7},
		{},
	}
	fmt.Println("All your base are belong to us!")
	// Starting moves
	if len(oppMoves) == 0 {
		return 5
	} else if len(oppMoves) == 1 {
		if oppMoves[0] != 5 && p.moveset[0] != 5 {
			return 5
		}
		for true {
			moves := []int{1, 3, 7, 9}
			rand.Seed(time.Now().UnixNano())
			randRes := rand.Intn(4)
			if isNewMove(moves[randRes], allMoves) {
				return moves[randRes]
			}
		}
	}
	//Finishing moves
	if len(p.moveset) >= 2 {
		xWin, xMove := comparator(xWinset, p.moveset, true, allMoves)
		yWin, yMove := comparator(yWinset, p.moveset, true, allMoves)
		dWin, dMove := comparator(dWinset, p.moveset, true, allMoves)
		if xWin {
			fmt.Println("ret x")
			fmt.Println(xMove)
			return xMove
		} else if yWin {
			fmt.Println("ret y")
			fmt.Println(yMove)
			return yMove
		} else if dWin {
			fmt.Println("ret d")
			fmt.Println(dMove)
			return dMove
		}
	}
	//
	// handle mid-game moves NOTE: TEMP, replace with defense?
	backup := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for true {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(9)
		move := backup[idx]
		fmt.Println(move)
		if isNewMove(move, allMoves) {
			fmt.Println("RANDOM MOVE")
			return move
		}
	}
	fmt.Println("RETURNING 0")
	return 0
}

// func isNewMove(move int, allMoves *[]int) bool {
// 	for _, idx := range *allMoves {
// 		if move == idx {
// 			return false
// 		}
// 	}
// 	return true
// }

// func inputHelper(playerName string) int {
// 	// Collect move input and convert to integer.
// 	var input string
// 	fmt.Println(playerName, "'s turn.")
// 	fmt.Scanln(&input)
// 	moveIndex, err := strconv.Atoi(input)
// 	if err != nil {
// 		for err != nil {
// 			fmt.Println("Please input a number from 0-9.")
// 			fmt.Scanln(&input)
// 			moveIndex, err = strconv.Atoi(input)
// 		}
// 		return moveIndex
// 	}
// 	return moveIndex
// }

// func comparator(winset [3][3]int, moveset []int, checkWin bool, allMoves *[]int) (bool, int) {
// 	// Returns true if moves match a single row in winset (3 in a row).
// 	for _, row := range winset {
// 		matchCnt := 0
// 		for _, val := range row {
// 			for _, move := range moveset {
// 				if val == move {
// 					matchCnt++
// 					if checkWin {
// 						if matchCnt == 2 {
// 							// if isNewMove(row[i+1], allMoves) {
// 							// 	return true, row[i+1]
// 							// }
// 							for _, blah := range row {
// 								if isNewMove(blah, allMoves) {
// 									return true, blah
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 		if matchCnt == 3 {
// 			return true, 0
// 		}
// 	}
// 	return false, 0
// }
