/* this is the code for five actions for Texas Holdem GameBoard
including Check(),Fold(),AllIn(),Call(),Raise(n int)
@Yichen Mo 0ct.16.2018
*/
package CoolProject

/*Check is the action that take the Player and update the attributes in the Player also update the Current
 */
func Check(player *Player) {
	(*player).OK = true
}

/*Fold is the action that player will hand out the holecards and will not participate the current board
 */
func Fold(currentBoard *Current, player *Player) {
	(*player).Fold = true
	numActivePlayers := 0
	for _, player := range currentBoard.Players{
		if !player.Eliminated && !player.Fold{
			numActivePlayers++
		}
	}
	if numActivePlayers == 1{
		for i, player := range currentBoard.Players{
			if !player.Eliminated && !player.Fold{
				(*currentBoard).Players[i].OK = true
			}
		}
	}

}

/*AllIn takes the player's all chips all his current chips to the pool
 */
func AllIn(currentBoard *Current, player *Player) {
	(*player).AllIn = true
	(*currentBoard).ChipPool = currentBoard.ChipPool + player.Chips
	(*player).Bet = player.Bet + player.Chips
	(*player).Chips = 0
	if currentBoard.CurrentBet < player.Bet {
		(*currentBoard).CurrentBet = player.Bet
		for i := range currentBoard.Players {
			(*currentBoard).Players[i].OK = false
		}
	}
	(*player).OK = true
}

/*Call take the input of the player and currentBoard, for this action the player would match the bet of current board
 */
func Call(currentBoard *Current, player *Player) {

	raisedAmount := currentBoard.CurrentBet - player.Bet
	(*player).Bet = currentBoard.CurrentBet
	(*player).Chips = player.Chips - raisedAmount
	(*currentBoard).ChipPool = currentBoard.ChipPool + raisedAmount
	(*player).OK = true
}

/*The player would raise the bet
 */
func Raise(currentBoard *Current, player *Player, numBet int) {
	for i := range currentBoard.Players {
		(*currentBoard).Players[i].OK = false
	}
	(*player).Bet = numBet + player.Bet
	(*currentBoard).CurrentBet = player.Bet
	(*currentBoard).ChipPool = currentBoard.ChipPool+numBet
	(*player).Chips = player.Chips - numBet
	(*player).OK = true

}
