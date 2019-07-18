package main

import "fmt"

func test() {
	allBoards := [][]int{}
	//enemy moves first, comp must compute all paths
	man := player{"X", []int{1, 2, 3, 4, 5, 6, 7}, 7, false, false, 1}
	_ = player{"O", []int{}, 0, true, false, 2}

	//create all boards w/ unique next moves
	availMoves := man.getAvailableMoves()
	for _, move := range availMoves {
		allBoards = append(allBoards, []int{move})
	}
	fmt.Println(allBoards)
}

func (p player) getAvailableMoves() (allMoves []int) {
	allMoves = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, pmove := range p.moves {
		for i, amove := range allMoves {
			if pmove == amove {
				allMoves[i] = allMoves[len(allMoves)-1]
				allMoves = allMoves[:len(allMoves)-1]
			}
		}
	}
	return
}
