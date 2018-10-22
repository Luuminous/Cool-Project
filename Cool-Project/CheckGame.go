/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
The function is to check if the game can continue. False means game end.
*/

func CheckGame(current Current) bool {
	numPlayers := 0 // To record how many players still alive.
	for _, player := range current.Players {
		if !player.Eliminated {
			numPlayers++
		}
	}
	if numPlayers <= 1 {
		return false
	}
	return true
}