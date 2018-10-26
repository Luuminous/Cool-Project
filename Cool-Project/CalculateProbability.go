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
