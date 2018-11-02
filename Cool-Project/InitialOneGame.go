/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
This function is to initiate everthing after one game.
*/

func InitialOneGame(current *Current){
	numPlayers := 0
	/*
	The loop is to reset several values from each uneliminated player and determine the start player.
	*/
	positionFinder := 2
	for index, player := range current.Players{
		if positionFinder == player.GamePosition {
			if player.Eliminated {
				positionFinder++
			} else {
				(*current).StartPlayer = player.SeatPosition
			}
		}
		if player.Eliminated == false {
			numPlayers++
			(*current).Players[index].OK = false
			(*current).Players[index].Fold = false
			(*current).Players[index].AllIn = false
			(*current).Players[index].Bet = 0
		}
	}
/*
	//Set the start position to UTG (now is small blind)
	count := 0
	start := false
	for (count <= 2) {
		for _, player := range current.Players {
			if start {
				if !player.Eliminated {
					count++
					if count == 2 {
						(*current).StartPlayer = player.SeatPosition
						break
					}
				}
			} else {
				if player.SeatPosition == current.StartPlayer {
					start = true
				}
			}
		}
	}
*/
	if numPlayers < 2 {
		panic("There aren't enough players! What's wrong with CheckGame()?")
	}
	/*
	Generate random sequence for the card.
	 */
	cardPool := Initiation()
	(*current).Pool = CreateRandomCards(2 * numPlayers + 5, cardPool)
	(*current).ChipPool = 0
	(*current).Stage = "Pre-flop"
	var temp []Card
	(*current).CommunityCard = temp
	var tempList []string
	(*current).PreEventsList = tempList
	/*
	Give hands to each player.
	*/
	count := 0
	for index := range current.Players {
		if current.Players[index].Eliminated == false {
			var temp []Card
			temp = append(temp, current.Pool[count])
			count++
			temp = append(temp, current.Pool[count])
			count++
			(*current).Players[index].Hands = temp
		}
	}
	/*
	The loop is to change the position of each uneliminated player.
	*/
	gamePosition := 1
	for (gamePosition <= numPlayers) {
		for index, player := range current.Players {
			if (gamePosition == 1) && (player.SeatPosition == current.StartPlayer) {
				(*current).Players[index].GamePosition = gamePosition
				if (*current).Players[index].Chips <= 10 {
					AllIn(current, &(*current).Players[index])
				} else {
					Raise(current, &(*current).Players[index], 10)
				}
				gamePosition++
				if gamePosition > numPlayers {
					break
				}
				continue
			}
			if (gamePosition > 1) && (player.Eliminated == false) {
				(*current).Players[index].GamePosition = gamePosition
				if gamePosition == 2 {
					if (*current).Players[index].Chips <= 20 {
						AllIn(current, &(*current).Players[index])
					} else {
						Raise(current, &(*current).Players[index], 20)
					}
					//Though Big Blind should raise 20, it still shouldn't be OK.
					(*current).Players[index].OK = false
				}
				gamePosition++
				if gamePosition > numPlayers {
					break
				}
				continue
			}
		}
	}
	InitialPlayerName(current)
}