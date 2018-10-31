/*
Author: Yichen Mo
Date: Oct. 31st, 2018
 */

package CoolProject

import "fmt"

func PrintRelatInfo(currentBoard *Current, player Player, playerActionList []string) {
	ClearScreen()
	fmt.Println("The current stage is " + currentBoard.Stage)
	fmt.Println(player.Name + ", the " + player.PositionName + ", now it's your turn")
	fmt.Println("Previous events: ")
	fmt.Println("--------------")
	for i := range currentBoard.PreEventsList {
		fmt.Println(currentBoard.PreEventsList[i])
	}
	fmt.Println("--------------")
	fmt.Println("Your Game position: ", player.GamePosition)
	fmt.Println("Your Holehands are " + HandsToString(player.Hands))
	fmt.Println("Your remaining chips are ", player.Chips)
	fmt.Println("Your have already put in wage", player.Bet)

	fmt.Println("Below is the information on the board")
	fmt.Println("The current Community Cards are " + HandsToString(currentBoard.CommunityCard))
	fmt.Println("The current Chip Pool is ", currentBoard.ChipPool)
	fmt.Println("The current Bet is ", currentBoard.CurrentBet)
	fmt.Println("The possible action lists in this round are ", playerActionList)

}