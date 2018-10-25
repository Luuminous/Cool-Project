package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/15/2018 1:20
*/

/*
ShowDown function is used to show the two cards of each player
The input is the reference of Current
The ouput is also the reference of Current
*/

func ShowDown(currentState *Current) []int {

	var playerSlice []Player

	for i := 0; i < len(currentState.Players); i++ {
		if !currentState.Players[i].Fold && !currentState.Players[i].Eliminated {
			playerSlice = append(playerSlice, currentState.Players[i])
		}
	}

	PrintShowDown(currentState)

	curMax := "00"
	var maxSeatNumber []int
	for i := 0; i < len(playerSlice); i++ {
		var sevenCardSlice []Card
		sevenCardSlice = append(sevenCardSlice, playerSlice[i].Hands[0], playerSlice[i].Hands[1])
		for j := 0; j < len(currentState.CommunityCard); j++ {
			sevenCardSlice = append(sevenCardSlice, currentState.CommunityCard[j])
		}
		compareString, _ := MaxPattern(sevenCardSlice)
		for m := 0; m < len(currentState.Players); m++ {
			if playerSlice[i].SeatPosition == currentState.Players[m].SeatPosition {
				(*currentState).Players[m].Pattern = compareString
			}
		}

		if compareString > curMax {
			curMax = compareString
		}
	}

	for i := 0; i < len(playerSlice); i++ {
		var sevenCardSlice []Card
		sevenCardSlice = append(sevenCardSlice, playerSlice[i].Hands[0], playerSlice[i].Hands[1])
		for j := 0; j < len(currentState.CommunityCard); j++ {
			sevenCardSlice = append(sevenCardSlice, currentState.CommunityCard[j])
		}
		compareString, _ := MaxPattern(sevenCardSlice)
		if curMax == compareString {
			maxSeatNumber = append(maxSeatNumber, playerSlice[i].SeatPosition)
		}
	}

	PrintWinner(currentState, maxSeatNumber)
	/*
		This loop is to set fold to all loser.
	*/
	for index, player := range currentState.Players {
		if IsIn(player.SeatPosition, maxSeatNumber) {
			(*currentState).Players[index].Fold = false
		} else {
			(*currentState).Players[index].Fold = true
		}
	}

	RepeatInput("yes")

	return maxSeatNumber
}
