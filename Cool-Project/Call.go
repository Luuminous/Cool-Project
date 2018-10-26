/* this is the code for five actions for Texas Holdem GameBoard
including Call()
@Yichen Mo 0ct.16.2018
*/
package CoolProject
/*Call take the input of the player and currentBoard, for this action the player would match the bet of current board
 */
func Call(currentBoard *Current, player *Player) {

	raisedAmount := currentBoard.CurrentBet - player.Bet
	(*player).Bet = currentBoard.CurrentBet
	(*player).Chips = player.Chips - raisedAmount
	(*currentBoard).ChipPool = currentBoard.ChipPool + raisedAmount
	(*player).OK = true

	(*currentBoard).PreEventsList = append((*currentBoard).PreEventsList, (player.Name + " chose to Call"))
}
