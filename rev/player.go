package main

import (
	"math/rand"
	"time"
)

type player struct {
	token string
	moves []int // possibly make fixed array||sized slice w/ make
	turns int
	isAI  bool
	win   bool
}

func generatePlayersRandomStart() (player, player) {
	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(2)
	token1, token2 := "X", "O"
	if rng == 1 {
		token1, token2 = token2, token1
	}
	p1 := player{token1, []int{}, 0, false, false}
	p2 := player{token2, []int{}, 0, false, false}
	if *aiFlag {
		rng2 := rand.Intn(2)
		p1.isAI = true
		if rng2 == 1 {
			p1.isAI, p2.isAI = false, true
		}
	}
	return p1, p2
}
