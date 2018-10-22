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
	handCards := len(currentState.Pool) - 5
	endOfFlop := len(currentState.Pool) - 2
	temp := currentState.Pool[handCards : endOfFlop]
	for _, card := range temp {
		(*currentState).CommunityCard = append(currentState.CommunityCard, card)
	}
}


/*
  Author: Chengyang Nie
  Date: 10/15/2018 12:58
*/

/*
Flop function is used to generate the turn card for the community []Cards.
The input of the function is the reference of Current.
*/

func Turn(currentState *Current) {
	//OutputTurn()
	
	turnCardIndex := len(currentState.Pool) - 2
	(*currentState).CommunityCard = append(currentState.CommunityCard, currentState.Pool[turnCardIndex])
}



/*
  Author: Chengyang Nie
  Date: 10/15/2018 1:04
*/

/*
Flop function is used to generate the river card for the community []Cards.
The input of the function is the reference of Current.
*/


func River(currentState *Current) {
	//OutputRiver()
	riverCardIndex := len(currentState.Pool) - 1
	(*currentState).CommunityCard = append(currentState.CommunityCard, currentState.Pool[riverCardIndex])
}