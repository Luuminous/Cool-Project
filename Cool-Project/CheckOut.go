/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
This function is to check out the state of the players and the game and split the chip pool.
*/

func CheckOut(current *Current) int {
	numPlayers := 0 // numPlayers check the number of players which is ok.
	for _, player := range current.Players {
		if !player.Eliminated && !player.Fold && (player.OK || player.AllIn) {
			numPlayers++
		}
	}
	if numPlayers == 1 {
		for index, player := range current.Players {
			if !player.Eliminated && !player.Fold && (player.OK || player.AllIn) {
				current.Players[index].Chips += current.ChipPool
			}
		}
		return 2
	}
	if numPlayers == 0 {
		panic("What's wrong with Command()?")
	}
	//numPlayers > 1
	for _, player := range current.Players {
		if !player.Eliminated && !player.Fold && !(player.OK || player.AllIn) {
			return 1
		}
	}
	return 0
}