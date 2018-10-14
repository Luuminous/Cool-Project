/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

/*
This function is calculate the max pattern from 7 cards.
*/
func MaxPattern(totalCards []Card) string {
	possibleCardsList := Generate5CardsFrom7Cards(totalCards)
	max := "00"
	for i := 0; i < len(possibleCardsList); i++ {
		temp := Convert5CardsToPatterns(possibleCardsList[i])
		if temp > max {
			max = temp
		}
	}
	return max
}