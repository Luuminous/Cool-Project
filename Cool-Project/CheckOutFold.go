/*
Author: Luuminous
Date: Oct. 31, 2018
 */

package CoolProject

func CheckOutFold(current *Current) {
	for index, player := range current.Players {
		if !player.Eliminated && !player.Fold{
			(*current).Players[index].Chips += current.ChipPool
			(*current).ChipPool = 0
		}
	}
}