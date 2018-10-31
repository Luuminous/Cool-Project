/* this is the code for the command for current round of Texas Holdem GameBoard

@Yichen 0ct.16.2018

*/
package CoolProject

import (
	"fmt"
	"strconv"
)

func Command(currentBoard *Current, startPosition int) {
	numActivePlayers := 0
	for _, player := range currentBoard.Players {
		if !player.Eliminated {
			numActivePlayers++
		}
	}
	currentPos := startPosition

	for currentPos <= numActivePlayers {
		for index, _ := range currentBoard.Players {
			if (*currentBoard).Players[index].GamePosition == currentPos {
				if (*currentBoard).Players[index].Eliminated {
					continue
				} else if (*currentBoard).Players[index].Fold {
					currentPos++
					continue
				} else if !(*currentBoard).Players[index].AllIn && !(*currentBoard).Players[index].OK {
					ClearScreen()
					fmt.Println("Please confirm your identity: " + currentBoard.Players[index].Name)
					RepeatInput("yes")
					//print the previous players' actions
					playerActionList := ChooseActions(currentBoard, (*currentBoard).Players[index])
					PrintRelatInfo(currentBoard, (*currentBoard).Players[index], playerActionList)
					action := InputInfo(playerActionList, (*currentBoard).Players[index].Chips,(*currentBoard).Players[index].Bet,*currentBoard.CurrentBet)
					if action[0] == 'R' {
						mon := action[6:]
						intMon, _ := strconv.Atoi(mon)
						Raise(currentBoard, &(*currentBoard).Players[index], intMon)
					} else if action == "Check" {
						Check(currentBoard, &(*currentBoard).Players[index])
					} else if action == "Call" {
						Call(currentBoard, &(*currentBoard).Players[index])
					} else if action == "AllIn" {
						AllIn(currentBoard, &(*currentBoard).Players[index])
					} else if action == "Fold" {
						Fold(currentBoard, &(*currentBoard).Players[index])
					}
					currentPos++
					fmt.Println("Your action has been executed, waiting for others...")
				} else {
					currentPos++
				}
			}

		}
	}
}


