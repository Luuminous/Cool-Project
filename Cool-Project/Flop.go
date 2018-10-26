package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/15/2018 12:45
*/

/*
Flop function is used to generate the three cards for the community []Cards.
The input of the function is the reference of Current.
*/

func Flop(currentState *Current) {
	//OutputFlop()
	(*currentState).Stage = "Flop"
	var tempList []string
	(*currentState).PreEventsList = tempList
	handCards := len(currentState.Pool) - 5
	endOfFlop := len(currentState.Pool) - 2
	temp := currentState.Pool[handCards:endOfFlop]
	for _, card := range temp {
		(*currentState).CommunityCard = append(currentState.CommunityCard, card)
	}
	for i := 0; i < len(currentState.Players); i++ {
		(*currentState).Players[i].OK = false
	}
}
