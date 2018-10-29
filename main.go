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
		winFlag = p1.collectPlayAndCheckWin(&allMoves)
		displayGame(p1, p2)
		if winFlag == false {
			winFlag = p2.collectPlayAndCheckWin(&allMoves)
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
		p1Char, p2Char = "O", "X"
	}
	p1 := player{[]int{}, 0, false, p1Char, false}
	p2 := player{[]int{}, 0, false, p2Char, false}
	if *aiFlag {
		anoRes := rand.Intn(2)
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

func (p *player) collectPlayAndCheckWin(allMoves *[]int) bool {
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
		moveIndex = p.aiGenerateNextMove(*allMoves, xWinset, yWinset, dWinset)
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
		dWin, _ := comparator(dWinset, p.moveset, false)
		xWin, _ := comparator(xWinset, p.moveset, false)
		yWin, _ := comparator(yWinset, p.moveset, false)
		if dWin || xWin || yWin {
			p.win = true
			return true
		} else if len(*allMoves) == 10 { // Tie game
			return true
		}
	}
	return false
}

func (ai *player) aiGenerateNextMove(allMoves []int, xSet, ySet, dSet [3][3]int) int {
	prepMoves := allMoves[1:]
	fmt.Println(prepMoves)
	// Check Win Condition
	res, matches := ai.CheckCanWin(xSet, ySet, dSet)
	fmt.Println(res, matches)
	switch res {
	case 0:
		rand.Seed(time.Now().UnixNano()) // NOTE: TEMP
		return rand.Intn(8) + 1
	case 1:
		// xwin
	case 2:
		// ywin
	case 3:
		// dwin
	}
	return 0
}

func (ai *player) CheckCanWin(xSet, ySet, dSet [3][3]int) (int, []int) {
	canWinX, xMtch := comparator(xSet, ai.moveset, true)
	canWinY, yMtch := comparator(ySet, ai.moveset, true)
	canWinD, dMtch := comparator(dSet, ai.moveset, true)
	if canWinX {
		return 1, xMtch
	} else if canWinY {
		return 2, yMtch
	} else if canWinD {
		return 3, dMtch
	} else {
		return 0, []int{}
	}
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

func comparator(winset [3][3]int, moveset []int, aiCheckFlag bool) (bool, []int) {
	// Returns true if moves match a single row in winset (3 in a row).
	for _, row := range winset {
		matchCnt := 0
		matches := []int{}
		for _, val := range row {
			for _, move := range moveset {
				if val == move {
					matchCnt++
					matches = append(matches, move)
				}
			}
		}
		if aiCheckFlag {
			if matchCnt == 2 {
				return true, matches
			}
		}
		if matchCnt == 3 {
			return true, matches
		}
	}
	return false, []int{}
}
