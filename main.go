package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type player struct {
	moveset []int
	turnCnt int
	win     bool
	char    string
	ai      bool
}

var aiFlag = flag.Bool("ai", false, "Computer opponent")

func main() {
	flag.Parse()
	allMoves := []int{0}
	indexMap := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, row := range indexMap {
		fmt.Println(row)
	}
	p1, p2 := generatePlayersRandomStart(aiFlag)
	winFlag := false
	for winFlag == false { // Turn management
		winFlag = p1.collectPlayAndCheckWin(&allMoves, p2.moveset)
		displayGame(p1, p2)
		if winFlag == false {
			winFlag = p2.collectPlayAndCheckWin(&allMoves, p1.moveset)
			displayGame(p1, p2)
		}
	}
	if p1.win {
		fmt.Println(p1.char, " Wins!")
	} else if p2.win {
		fmt.Println(p2.char, " Wins!")
	} else {
		fmt.Println("Tie Game!")
	}
	// For compiled code, to avoid terminal auto-close
	var null string
	fmt.Scanln(&null)
}

func generatePlayersRandomStart(aiFlag *bool) (player, player) {
	rand.Seed(time.Now().UnixNano())
	randRes := rand.Intn(2)
	p1Char, p2Char := "X", "O"
	if randRes == 1 {
		p1Char, p2Char = "O", "X" // NOTE: TEMP CMMT
	}
	// p1 := player{[]int{1, 3, 5}, 0, false, p1Char, false} // NOTE: sim near-win
	// p2 := player{[]int{2, 6, 7}, 0, false, p2Char, false} //
	p1 := player{[]int{}, 0, false, p1Char, false}
	p2 := player{[]int{}, 0, false, p2Char, false}
	if *aiFlag {
		// p2.ai = true
		anoRes := rand.Intn(2) // NOTE: TEMP CMMT
		if anoRes == 1 {
			p1.ai = true
		} else {
			p2.ai = true
		}
	}
	return p1, p2
}

func displayGame(p1 player, p2 player) {
	charset := [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	displayHelper(&charset, p1.moveset, p1.char)
	displayHelper(&charset, p2.moveset, p2.char)
	for _, row := range charset {
		fmt.Println(row)
	}
}

func displayHelper(charset *[3][3]string, moveset []int, playChar string) {
	// Place string characters ('X', 'O') on each players move index.
	for _, move := range moveset {
		if move <= 3 {
			charset[0][move-1] = playChar
		} else if move <= 6 {
			charset[1][move-4] = playChar
		} else {
			charset[2][move-7] = playChar
		}
	}
}

func (p *player) collectPlayAndCheckWin(allMoves *[]int, oppMoves []int) bool {
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
	moveIndex := 0
	if p.ai {
		// NOTE: TEMP
		//
		//
		fmt.Println("All Moves: ", *allMoves)
		fmt.Println("My moves: ", p.moveset)
		fmt.Println("Your moves: ", oppMoves)
		rand.Seed(time.Now().UnixNano())
		moveIndex = rand.Intn(9) + 1
		fmt.Println("Made move: ", moveIndex)
	} else { // Human input
		passCnt := 0
		moveIndex = inputHelper(p.char)
		for passCnt != len(*allMoves) { // Loop until valid move played
			for _, move := range *allMoves {
				if move == moveIndex {
					fmt.Println("Move already played.")
					passCnt = 0
					moveIndex = inputHelper(p.char)
					break
				} else {
					// Limit move input 0-9
					if moveIndex <= 9 {
						passCnt++
					} else {
						moveIndex = 0
					}
				}
			}
		}
	}
	*allMoves = append(*allMoves, moveIndex)
	p.moveset = append(p.moveset, moveIndex)
	p.turnCnt++
	if p.turnCnt > 2 {
		dWin := comparator(dWinset, p.moveset)
		xWin := comparator(xWinset, p.moveset)
		yWin := comparator(yWinset, p.moveset)
		if dWin || xWin || yWin {
			p.win = true
			return true
		} else if len(*allMoves) == 10 { // Tie game
			return true
		}
	}
	return false
}

func test() {
	// e. players moves
	// available moves/spaces
	// check win state
}

func inputHelper(playerName string) int {
	// Collect move input and convert to integer.
	var input string
	fmt.Println(playerName, "'s turn.")
	fmt.Scanln(&input)
	moveIndex, err := strconv.Atoi(input)
	if err != nil {
		for err != nil {
			fmt.Println("Please input a number from 0-9.")
			fmt.Scanln(&input)
			moveIndex, err = strconv.Atoi(input)
		}
		return moveIndex
	}
	return moveIndex
}

func comparator(winset [3][3]int, moveset []int) bool {
	// Returns true if moves match a single row in winset (3 in a row).
	for _, row := range winset {
		matchCnt := 0
		for _, val := range row {
			for _, move := range moveset {
				if val == move {
					matchCnt++
				}
			}
		}
		if matchCnt == 3 {
			return true
		}
	}
	return false
}
