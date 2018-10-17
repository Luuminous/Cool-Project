package CoolProject

import "fmt"

/*
  Author: Chengyang Nie
  Date: 10/15/2018 1:20
*/


/*
ShowDown function is used to show the two cards of each player
The input is the reference of Current
The ouput is also the reference of Current
*/

func ShowDown(currentState *Current) {

	var playerSlice []Player

	for i := 0; i < len(currentState.Players); i++ {
		if !currentState.Players[i].Fold {
			playerSlice = append(playerSlice, currentState.Players[i])
		}
	}

	PrintShowDown(currentState)

	curMax = "00"
	var finalSlice []Card
	var maxSeatNumber int
	for i := 0; i < len(playerSlice); i++ {
		var 7CardSlice []Card
		7CardSlice = append(7CardSlice, playerSlice[i].Hands[0], playerSlice[i].Hands[1])
		for j := 0; j < len(currentState.CommunityCard); j++ {
			7CardSlice = append(7CardSlice, currentState.CommunityCard[j])
		}


		compareString, sliceCard := MaxPattern(7CardSlice)
		if compareString > curMax {
			finalSlice = sliceCard
			curMax = compareString
			maxSeatNumber = playerSlice[i].SeatPosition
		}
	}


	PrintWinner(currentState, maxSeatNumber, finalSlice)

}


func PrintShowDown(currentState *Current) {
	var playerSlice []Player

	for i := 0; i < len(currentState.Players); i++ {
		if !currentState.Players[i].Fold {
			playerSlice = append(playerSlice, currentState.Players[i])
		}
	}
	for i := 0; i < len(playerSlice); i++ {
		fmt.Println("The player at seat number " + playerSlice.Player[i].SeatPosition + " has " + playerSlice.Player[i].Hands[0] + " " + playerSlice.Player[i].Hands[1])
	}
}

func PrintWinner(currentState *Current, maxSeatNumber int, finalSlice []Card) {
	fmt.Println("The winner is the palyer at seat number " + maxSeatNumber + " and the max pattern is " + HandsToString(finalSlice))
}
