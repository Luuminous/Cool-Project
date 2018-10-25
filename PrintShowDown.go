/*PrintShowDown will print the ShowDown results for player @Chengyang*/
package CoolProject

import (
	"fmt"
)

func PrintShowDown(currentState *Current) {
	var playerSlice []Player

	for i := 0; i < len(currentState.Players); i++ {
		if !currentState.Players[i].Fold && !currentState.Players[i].Eliminated {
			playerSlice = append(playerSlice, currentState.Players[i])
		}
	}
	for i := 0; i < len(playerSlice); i++ {
		fmt.Print("The player ")
		fmt.Print(playerSlice[i].Name)
		fmt.Println(" has " + HandsToString(playerSlice[i].Hands))
	}

	fmt.Println("The CommunityCards are " + HandsToString(currentState.CommunityCard))
}
