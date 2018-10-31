/*
Author: Luuminous
Date: Oct. 30th, 2018
 */

package CoolProject

import (
	"fmt"
)

func PrintFoldResult(current *Current) {
	winnerIndex := 0
	for index, player := range current.Players {
		if !player.Eliminated && !player.Fold {
			winnerIndex = index
		}
	}
	fmt.Println("The Chip Pool now is: ", current.ChipPool)
	fmt.Println(current.Players[winnerIndex].Name + " wins all the chips!!")
	RepeatInput("yes")
}