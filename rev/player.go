package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type player struct {
	token string
	moves []int
	turns int
	isAI  bool
	win   bool
}

func generatePlayersRandomStart() (p1, p2 player) {
	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(2)
	token1, token2 := "X", "O"
	if rng == 1 {
		token1, token2 = token2, token1
	}
	p2 = player{"X", []int{3, 5}, 2, true, false} // TEMP, !SWAPPED P VARS!
	p1 = player{"O", []int{}, 0, false, false}
	// if *aiFlag {
	// 	rng2 := rand.Intn(2)
	// 	p1.isAI = true
	// 	if rng2 == 1 {
	// 		p1.isAI, p2.isAI = false, true
	// 	}
	// }
	return
}

func (p *player) collectPlay() (win bool) {
	moveIndex := 0
	if p.isAI {
		moveIndex = p.generatePlay()
		fmt.Println("Comp made move: ", moveIndex)
	} else {
		validatingMove := true
		moveIndex = p.inputHelper()
		for validatingMove {
			if moveIndex > 9 {
				moveIndex = 0
			}
			if isNewMove(moveIndex) {
				validatingMove = false
			} else {
				fmt.Println("Move invalid or already played.")
				moveIndex = p.inputHelper()
			}
		}
	}
	allPlayedMoves = append(allPlayedMoves, moveIndex)
	p.moves = append(p.moves, moveIndex)
	p.turns++
	win = false
	if p.turns > 2 {
		win = p.checkWin()
	}
	return
}

func (p player) generatePlay() int {
	// 1. Check for self win, ret win move
	res, move := comparator(p.moves, true)
	if res {
		return move
	}
	// 2. Check for opp win, ret win move (block)
	// 3. Attempt self fork(2 possible wins), ret move
	// 4. Return center(5)
	// 5. Ret corner opposite to opp
	// 6. Ret random
	return 0
}

func (p player) inputHelper() (moveIndex int) {
	var input string
	fmt.Println(p.token, "'s turn.")
	fmt.Scanln(&input)
	moveIndex, err := strconv.Atoi(input)
	for err != nil {
		fmt.Println("Please input a number from 0-9.")
		fmt.Scanln(&input)
		moveIndex, err = strconv.Atoi(input)
	}
	return
}

func (p *player) checkWin() bool {
	win, _ := comparator(p.moves, false)
	if win {
		p.win = true
	} else if len(allPlayedMoves) == 10 {
		return true
	}
	return win
}
