/*
Author: Yichen Mo
Date: Oct. 31st, 2018
 */

package CoolProject


func ChooseActions(currentBoard *Current, player Player) []string {
	actionList := []string{}
	raiseMoney := currentBoard.CurrentBet - player.Bet
	if raiseMoney == 0 {
		//check raise
		actionList = append(actionList, "Check", "Raise", "AllIn")
	} else {
		if raiseMoney >= player.Chips {
			actionList = append(actionList, "AllIn", "Fold")
		} else {
			actionList = append(actionList, "Call", "AllIn", "Fold", "Raise")
		}

	}
	return actionList

}
