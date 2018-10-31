/*
Author: Luuminous
Date: Oct. 31st, 2018
 */

package CoolProject

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

/*Take Input information and execute the actions */
func InputInfo(actionList []string, max int) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your action: ")
	action, _, _ := reader.ReadLine()
	aStr := string(action)
	for !IsInString(aStr, actionList) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Wrong input, enter your action again: ")
		action, _, _ := reader.ReadLine()
		aStr = string(action)
	}
	if aStr == "Raise" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter how much money you want to put: ")
		money, _, _ := reader.ReadLine()
		aMoney := string(money)
		intMon, err := strconv.Atoi(aMoney)

		for (intMon >= max) || err != nil {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Wrong money number, please enter again: ")
			money, _, _ = reader.ReadLine()
			aMoney = string(money)
			intMon, err = strconv.Atoi(aMoney)
		}
		aStr = aStr + " " + aMoney
	}
	fmt.Println("Executing your action : " + aStr)
	RepeatInput("yes")
	return aStr
}

