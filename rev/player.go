package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type player struct {
	token string
	moves []int // possibly make fixed array||sized slice w/ make
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
	p1 = player{token1, []int{}, 0, false, false}
	p2 = player{token2, []int{}, 0, false, false}
	if *aiFlag {
		rng2 := rand.Intn(2)
		p1.isAI = true
		if rng2 == 1 {
			p1.isAI, p2.isAI = false, true
		}
	}
	return
}

func (p *player) collectPlay() bool {
	moveIndex := 0
	if p.isAI {
		//
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
	if p.turns > 2 {
		return p.checkWin()
	}
	return false
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
