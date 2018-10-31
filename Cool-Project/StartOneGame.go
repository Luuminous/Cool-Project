package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/14/2018 11:49
*/

/*
This function is used to start one Texas Holdem game.
The input of the function is the Current Consturct.
This function will contain subroutines of one game.

*/

func StartOneGame(currentState *Current) {
	numPlayers := 0
	for _, player := range currentState.Players {
		if !player.Eliminated {
			numPlayers++
		}
	}
	InitialOneGame(currentState)
	isFirst := true
	for CheckOut(currentState) == 1 {
		if isFirst{
			if numPlayers >= 3 {
				Command(currentState, 3)
				isFirst = false
			} else {
				Command(currentState, 1)
				isFirst = false
			}
		} else {
			Command(currentState, 1)
		}
	}
	if CheckOut(currentState) == 2 {
		PrintFoldResult(currentState)
		CheckOutFold(currentState)
		(*currentState).GameCount++
		return
	}
	Flop(currentState)
	for CheckOut(currentState) == 1 {
		Command(currentState, 1)
	}
	if CheckOut(currentState) == 2 {
		PrintFoldResult(currentState)
		CheckOutFold(currentState)
		(*currentState).GameCount++
		return
	}
	Turn(currentState)
	for CheckOut(currentState) == 1 {
		Command(currentState, 1)
	}
	if CheckOut(currentState) == 2 {
		PrintFoldResult(currentState)
		CheckOutFold(currentState)
		(*currentState).GameCount++
		return
	}
	River(currentState)
	for CheckOut(currentState) == 1 {
		Command(currentState, 1)
	}
	if CheckOut(currentState) == 2 {
		PrintFoldResult(currentState)
		CheckOutFold(currentState)
		(*currentState).GameCount++
		return
	}

	winner := ShowDown(currentState)

	FinalCheckOut(currentState, winner)
	(*currentState).GameCount++

}
