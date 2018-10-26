/* this is the code for five actions for Texas Holdem GameBoard
including Fold()
@Yichen Mo 0ct.16.2018
*/
package CoolProject

/*Fold is the action that player will hand out the holecards and will not participate the current board
 */
func Fold(currentBoard *Current, player *Player) {
	(*player).Fold = true
	(*currentBoard).PreEventsList = append((*currentBoard).PreEventsList, (player.Name + " chose to fold"))
	numActivePlayers := 0
	for _, player := range currentBoard.Players {
		if !player.Eliminated && !player.Fold {
			numActivePlayers++
		}
	}
	if numActivePlayers == 1 {
		for i, player := range currentBoard.Players {
			if !player.Eliminated && !player.Fold {
				(*currentBoard).Players[i].OK = true
			}
		}
	}

}
