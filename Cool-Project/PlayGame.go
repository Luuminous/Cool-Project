/*
Author: Luuminous
Date: Oct. 17th, 2018
*/
package CoolProject

import "fmt"

type Player struct {
	Hands []Card
	SeatPosition int
	GamePosition int
	Chips, Bet int
	Fold, AllIn, OK, Eliminated bool
	PositionName string
}

type Current struct {
	Pool []Card
	CommunityCard []Card
	ChipPool int
	Players []Player
	StartPlayer int
	CurrentBet int
	CurrentPlayerPosition int
	GameCount int
}

/*
This function is to play the Texas Holdem game.
*/
func PlayGame(numPlayers, initialChips, numTurns int) {
	if numPlayers <= 1 {
		panic("Too little players!")
	}
	if initialChips <= 100 {
		panic("Too little chips!")
	}
	if numTurns < 1 {
		panic("Too little turns!")
	}
	var current Current = InitialGame(numPlayers, initialChips)

	for (current.GameCount <= numTurns) {
		StartOneGame(&current)
	}
}