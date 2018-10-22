/* this is the code for the command for current round of Texas Holdem GameBoard

@Yichen 0ct.16.2018

*/
package CoolProject

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Command(currentBoard *Current) {
	numActivePlayers := 0
	for _, player := range currentBoard.Players {
		if !player.Eliminated {
			numActivePlayers++
		}
	}
	currentPos := 1
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
					playerActionList := ChooseActions(currentBoard, (*currentBoard).Players[index])
					PrintRelatInfo(currentBoard, (*currentBoard).Players[index], playerActionList)
					action := InputInfo(playerActionList, (*currentBoard).Players[index].Chips)
					if action[0] == 'R' {
						mon := action[6:]
						intMon, _ := strconv.Atoi(mon)
						Raise(currentBoard, &(*currentBoard).Players[index], intMon)
					} else if action == "Check" {
						Check(&(*currentBoard).Players[index])
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

func ChooseActions(currentBoard *Current, player Player) []string {
	actionList := []string{}
	raiseMoney := currentBoard.CurrentBet - player.Bet
	if raiseMoney == 0 {
		//check raise
		actionList = append(actionList, "Check", "Raise", "AllIn")
	} else {
		if raiseMoney >= player.Chips {
			actionList = append(actionList, "AllIn", "Fold")
		} else {
			actionList = append(actionList, "Call", "AllIn", "Fold", "Raise")
		}

	}
	return actionList

}

func PrintRelatInfo(currentBoard *Current, player Player, playerActionList []string) {
	ClearScreen()
	fmt.Println(player.Name + ", now it's your turn")
	fmt.Println("The current stage is "+ currentBoard.Stage)
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
func ClearScreen(){
	//clean the screen
	for i := 0; i < 56; i++ {
		fmt.Println()
	}
}

/*Take Input information and execute the actions */
func InputInfo(actionList []string, max int) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your action: ")
	action, _ := reader.ReadString('\n')
	action = action[:len(action)-1]
	for !IsInString(action, actionList) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Wrong input, enter your action again: ")
		action, _ = reader.ReadString('\n')
		action = action[:len(action)-1]
	}
	if action == "Raise" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter how much money you want to put: ")
		money, _ := reader.ReadString('\n')
		money = money[:len(money)-1]
		intMon, err := strconv.Atoi(money)

		for ((intMon >= max) || err != nil) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Wrong money number, please enter again: ")
			money, _ := reader.ReadString('\n')
			money = money[:len(money)-1]
			intMon, err = strconv.Atoi(money)
		}
		action = action + " " + money
	}
	fmt.Println("Executing your action : " + action)
	RepeatInput("yes")
	return action
}

func IsInString(s string, array []string) bool {
	for _, val := range array {
		if val == s {
			return true
		}
	}
	return false
}

/**/
