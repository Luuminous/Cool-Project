/*
Author: Luuminous
Date: Oct. 30, 2018
*/

package CoolProject

import (
	"fmt"
)

func PrintGameResult(current Current) {
	for _, player := range current.Players {
		fmt.Println(player.SeatPosition, " ", player.Name, " ", player.Chips)
	}
	fmt.Println("Happy everyday!!!")
}