/*
  Author: Chengyang Nie
  Date: 10/15/2018 12:58
*/
package CoolProject
/*
Flop function is used to generate the turn card for the community []Cards.
The input of the function is the reference of Current.
*/

func Turn(currentState *Current) {
	//OutputTurn()
	(*currentState).Stage = "Turn"
	var tempList []string
	(*currentState).PreEventsList = tempList
	turnCardIndex := len(currentState.Pool) - 2
	(*currentState).CommunityCard = append(currentState.CommunityCard, currentState.Pool[turnCardIndex])

	for i := 0; i < len(currentState.Players); i++ {
		(*currentState).Players[i].OK = false
	}
}
