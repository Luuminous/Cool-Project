/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject

/*
This function is to initiate everthing after one game.
*/

func InitialOneGame(current *Current) {
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
				current.StartPlayer = player.SeatPosition
			}
		}
		if player.Eliminated == false {
			numPlayers++
			current.Players[index].OK = false
			current.Players[index].Fold = false
		}
	}
	if numPlayers < 2 {
		panic("There aren't enough players! What's wrong with Check()?")
	}
	/*
	The loop is to change the position of each uneliminated player.
	 */
	gamePosition := 1
	for (gamePosition <= numPlayers) {
		for index, player := range current.Players {
			if (gamePosition == 1) && (player.SeatPosition == current.StartPlayer) {
				current.Players[index].GamePosition = gamePosition
				gamePosition++
				if gamePosition > numPlayers {
					break
				}
				continue
			}
			if (gamePosition > 1) && (player.Eliminated == false) {
				current.Players[index].GamePosition = gamePosition
				gamePosition++
				if gamePosition > numPlayers {
					break
				}
				continue
			}
		}
	}
	InitialPlayerName(current)
	/*
	Generate random sequence for the card.
	 */
	cardPool := Initiation()
	current.Pool = CreateRandomCards(2 * numPlayers + 5, cardPool)
}