
/*
  Author: Chengyang Nie
  Date: 10/15/2018 1:04
*/
package CoolProject
/*
Flop function is used to generate the river card for the community []Cards.
The input of the function is the reference of Current.
*/


func River(currentState *Current) {
	//OutputRiver()
	(*currentState).Stage = "River"
	var tempList []string
	(*currentState).PreEventsList = tempList
	riverCardIndex := len(currentState.Pool) - 1
	(*currentState).CommunityCard = append(currentState.CommunityCard, currentState.Pool[riverCardIndex])

	for i := 0; i < len(currentState.Players); i++ {
		(*currentState).Players[i].OK = false
	}
}
