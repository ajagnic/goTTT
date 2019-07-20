package main

import "fmt"

type board struct {
	moves []int
	id    int
}

var counter int

func test() {
	counter = 0
	man := player{"X", []int{}, 0, false, false, 1}
	comp := player{"O", []int{}, 0, true, false, 2}
	manOpenMoves := man.getAvailableMoves()
	fmt.Println(manOpenMoves)
	manPossibleGames := generateBoardsForMoves(manOpenMoves)
	fmt.Println(manPossibleGames)
	for _, game := range manPossibleGames {
		compOpenMoves := comp.getAvailableMoves()
		fmt.Println(compOpenMoves)
		compPossibleGames := generateBoardsForMoves(compOpenMoves)
		fmt.Println(compPossibleGames)
	}

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

func generateBoardsForMoves(openMoves []int) (allBoards []board) {
	for _, move := range openMoves {
		moves := []int{move}
		newBoard := board{moves, counter}
		counter++
		allBoards = append(allBoards, newBoard)
	}
	return
}

func removeFromSlc() {}
