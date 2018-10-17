/*
Author: Luuminous
Date: Oct. 17th, 2018
*/

package CoolProject


/*
This function is to initiate the name according to each player's position.
*/

/*
position 1: Small Blind
position 2: Big Blind
position 3: Under the Gun
position 4: Middle Position
.
.
.
position -2: Cut Off
position -1: Button
*/
		
func InitialPlayerName(current *Current) {
	numPlayers := 0
	for _, player := range current.Players {
		if player.Eliminated == false {
			numPlayers++
		}
	}
	nameList := make([]string, numPlayers)
	nameList[0] = "Small Blind"
	nameList[1] = "Big Blind"
	if numPlayers == 3 {
		nameList[2] = "Button"
	} else if numPlayers == 4 {
		nameList[2] = "Under the Gun"
		nameList[3] = "Button"
	} else if numPlayers == 5 {
		nameList[2] = "Under the Gun"
		nameList[3] = "Cut Off"
		nameList[4] = "Button"
	} else if numPlayers >= 6{
		nameList[2] = "Under the Gun"
		for i := 3; i < numPlayers - 2; i++ {
			nameList[i] = "Middle Position"
		}
		nameList[numPlayers - 2] = "Cut Off"
		nameList[numPlayers - 1] = "Button"
	}
	for index, player := range current.Players {
		if player.Eliminated == false {
			current.Players[index].PositionName = nameList[player.GamePosition - 1]
		}
	}
}