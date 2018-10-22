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
	numPlayers := 0 // numPlayers check the number of players which isn't fold.
	for _, player := range current.Players {
		if !player.Eliminated && !player.Fold {
			numPlayers++
		}
	}
	if numPlayers == 1 {
		for index, player := range current.Players {
			if !player.Eliminated && !player.Fold{
				winnerIndex := index
				(*current).Players[index].Chips += current.ChipPool
			}
		}
		PrintFoldResult(current, winnerIndex)
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

func PrintFoldResult(current *Current, index int) {
	fmt.Println("The Chip Pool now is: ", current.ChipPool)
	fmt.Println(current.Players[index].Name + " wins all the chips!!")
	RepeatInput("yes")
}