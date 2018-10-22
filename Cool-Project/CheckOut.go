/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
This function is to check out the state of the players and the game and split the chip pool.
*/

func CheckOut(current *Current) int {
	activePlayers := 0 // Record the active players.
	for _, player := range current.Players {
		if !player.Eliminated{
			activePlayers++
		}
	}
	numPlayers := 0 // numPlayers check the number of players which is fold.
	for _, player := range current.Players {
		if !player.Eliminated && player.Fold {
			numPlayers++
		}
	}
	if activePlayers - numPlayers == 1 {
		for index, player := range current.Players {
			if !player.Eliminated && player.Fold{
				(*current).Players[index].Chips += current.ChipPool
			}
		}
		return 2
	}
	// More than 1 player not fold.
	for _, player := range current.Players {
		if !player.Eliminated && !player.Fold && !(player.OK || player.AllIn) {
			return 1
		}
	}
	return 0
}