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
	InitialOneGame(currentState)
	for (CheckOut(currentState) == 1) {
		Command(currentState)
	}
	if CheckOut(currentState) == 2 {
		(*currentState).GameCount++
		return
	}
	Flop(currentState)
	for (CheckOut(currentState) == 1) {
		Command(currentState)
	}
	if CheckOut(currentState) == 2 {
		(*currentState).GameCount++
		return
	}
	Turn(currentState)
	for (CheckOut(currentState) == 1) {
		Command(currentState)
	}
	if CheckOut(currentState) == 2 {
		(*currentState).GameCount++
		return
	}
	River(currentState)
	for (CheckOut(currentState) == 1) {
		Command(currentState)
	}
	if CheckOut(currentState) == 2 {
		(*currentState).GameCount++
		return
	}

	winner := ShowDown(currentState)

	FinalCheckOut(currentState, winner)
	(*currentState).GameCount++

}












