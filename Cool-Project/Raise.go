/* this is the code for five actions for Texas Holdem GameBoard
including,Raise(n int)
@Yichen Mo 0ct.16.2018
*/
package CoolProject

import "strconv"

/*The player would raise the bet
 */
func Raise(currentBoard *Current, player *Player, numBet int) {
	for i := range currentBoard.Players {
		(*currentBoard).Players[i].OK = false
	}
	(*player).Bet = numBet + player.Bet
	(*currentBoard).CurrentBet = player.Bet
	(*currentBoard).ChipPool = currentBoard.ChipPool + numBet
	(*player).Chips = player.Chips - numBet
	(*player).OK = true

	stringInput := player.Name + " chose to Raise "
	stringInput = stringInput + strconv.Itoa(numBet)
	(*currentBoard).PreEventsList = append((*currentBoard).PreEventsList, stringInput)

}
