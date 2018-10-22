/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
The function is for the finally check out after showdown.
*/

func IsIn(element int, set []int) bool {
	for _, value := range set {
		if value == element {
			return true
		}
	}
	return false
}

func FinalCheckOut(current *Current, winner []int) {
	//Make the all-in loser eliminated.
	for index,player := range current.Players {
		if !player.Eliminated && player.AllIn && !IsIn(player.SeatPosition, winner){
			(*current).Players[index].Eliminated = true
		}
	}
	//Calculate how many chips each players can win.
	totalBetAmount := 0
	for _, player := range current.Players {
		if IsIn(player.SeatPosition, winner) {
			totalBetAmount += player.Bet
		}
	}
	for index, player := range current.Players {
		if IsIn(player.SeatPosition, winner) {
			(*current).Players[index].Chips += current.ChipPool * player.Bet / totalBetAmount
		}
	}
}