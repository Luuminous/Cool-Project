/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

import "math/rand"

/*
This function is to calculate the possibility for a given hand. 
*/
func Simulation(hands []Card, numTrials int) float64 {
	count := 0
	initiationPool := Initiation()
	afterDrawPool := DeleteFormPool(initiationPool, hands)
	for i := 0; i < numTrials; i++ {
		sevenPool := CreateRandomCards(7, afterDrawPool)
		first := rand.Intn(7)
		second := rand.Intn(7)
		for (first == second) {
			second = rand.Intn(7)
		}
		yourPool := hands
		for i := 0; i < 7; i++ {
			if (i != first) && (i != second) {
				yourPool = append(yourPool, sevenPool[i])
			}
		}
		oppoPattern := MaxPattern(sevenPool)
		yourPattern := MaxPattern(yourPool)
		if yourPattern > oppoPattern {
			count++
		}
	}
	return float64(count) / float64(numTrials)
}