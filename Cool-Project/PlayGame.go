/*
Author: Luuminous
Date: Oct. 17th, 2018
*/
package CoolProject

type Player struct {
	Hands []Card
	SeatPosition int
	GamePosition int
	Chips, Bet int
	Fold, AllIn, OK, Eliminated bool
	PositionName string
	Pattern string
}

type Current struct {
	Pool []Card
	CommunityCard []Card
	ChipPool int
	Players []Player
	StartPlayer int
	CurrentBet int
	GameCount int
	ActivePlayers []Player
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

	for (current.GameCount <= numTurns && CheckGame(current)) {
		StartOneGame(&current)
	}
}