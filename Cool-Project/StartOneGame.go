package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/14/2018 11:49
*/

/*
This function is used to start one Texas Holdem game.
The input of the function is the Current Insturct.
The output is also the Current Instruct.
This function will contain subroutines of one game.

*/

func StartOneGame(currentState *Current) {
	//shuffle 52 cards
	InitialOneGame(&currentState)

	Hands(&currentState)
	Command(&currentState)

	if CheckOut(&currentState) {

		Flop(&currentState)
		Command(&currentState)

		if CheckOut(&currentState) {

			Turn(&currentState)
			Command(&currentState)

			if CheckOut(&currentState) {

				River(&currentState)
				Command(&currentState)

				if CheckOut(&currentState) {

					ShowDown(&currentState)
					// Check is used to check the number of players
					CheckOut(&currentState)
				}
			}
		}
	}
	
}












