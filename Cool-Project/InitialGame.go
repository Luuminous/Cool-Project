/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
This function is to initiate the whole game.
*/
func InitialGame(numPlayers, initialChips int) Current {
	var newCurrent Current
	newCurrent.ChipPool = 0
	var newPlayerList []Player
	for i := 1; i <= numPlayers; i++ {
		var newPlayer Player
		newPlayer.SeatPosition = i
		newPlayer.GamePosition = i
		newPlayer.Chips = initialChips
		newPlayer.Bet = 0
		newPlayer.Fold = false
		newPlayer.AllIn = false
		newPlayer.OK = false
		newPlayer.Eliminated = false
		newPlayerList = append(newPlayerList, newPlayer)
	}
	newCurrent.Players = newPlayerList
	newCurrent.StartPlayer = 1
	newCurrent.CurrentBet = 0
	var nilCard []Card
	newCurrent.CommunityCard = nilCard
	newCurrent.Pool = nilCard
	InitialPlayerName(&newCurrent)
	return newCurrent
}