/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

import (
	"fmt"
	"os"
	"strconv"
)

/*
The function is to convert the hands to a readable string.
*/
func HandsToString(hands []Card) string {
	ans := ""
	for i := 0; i < len(hands); i++ {
		num := hands[i].Num
		if num == 14 {
			ans += "A"
		} else if num == 13 {
			ans += "K"
		} else if num == 12 {
			ans += "Q"
		} else if num == 11 {
			ans += "J"
		} else {
			ans += strconv.Itoa(num)
		}
		color := hands[i].Color
		if color == 0 {
			ans += "♠"
		} else if color == 1 {
			ans += "♥"
		} else if color == 2 {
			ans += "♣"
		} else if color == 3 {
			ans += "♦"
		}
		ans += " "
	}
	return ans
}

/*
The function is to calculate all the probability and record them into a file.
*/
func CalculateProbability(numTrials int) {
	fileName := strconv.Itoa(numTrials) + ".txt"
	file, _ := os.Create(fileName)
	pool := Initiation()
	count := 0
	for i := 0; i < len(pool)-1; i++ {
		for j := i + 1; j < len(pool); j++ {
			var hands []Card
			hands = append(hands, pool[i], pool[j])
			prob := Simulation(hands, numTrials)
			file.WriteString(HandsToString(hands) + "\t" + strconv.FormatFloat(prob, 'f', 4, 64) + "\n")
			fmt.Println(count)
			count++
		}
	}
}
