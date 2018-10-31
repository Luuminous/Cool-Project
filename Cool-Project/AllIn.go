/* this is the code for five actions for Texas Holdem GameBoard
including AllIn()
@Yichen Mo 0ct.16.2018
*/
package CoolProject
/*AllIn takes the player's all chips all his current chips to the pool
 */
func AllIn(currentBoard *Current, player *Player) {
	(*player).AllIn = true
	(*currentBoard).ChipPool = currentBoard.ChipPool + player.Chips
	(*player).Bet = player.Bet + player.Chips
	(*player).Chips = 0
	if player.Bet > currentBoard.CurrentBet {
		(*currentBoard).CurrentBet = player.Bet
		for index := range currentBoard.Players {
			if currentBoard.Players[index].SeatPosition != player.SeatPosition && !currentBoard.Players[index].Eliminated && !currentBoard.Players[index].Fold {
				(*currentBoard).Players[index].OK = false
			}
		}

	}
	(*currentBoard).PreEventsList = append((*currentBoard).PreEventsList, (player.Name + " chose to AllIn"))
	if currentBoard.CurrentBet < player.Bet {
		(*currentBoard).CurrentBet = player.Bet
		for i := range currentBoard.Players {
			(*currentBoard).Players[i].OK = false
		}
	}
	(*player).OK = true

}
