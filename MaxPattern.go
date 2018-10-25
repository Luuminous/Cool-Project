/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

/*
This function is calculate the max pattern from 7 cards.
*/
func MaxPattern(totalCards []Card) (string, []Card) {
	possibleCardsList := Generate5CardsFrom7Cards(totalCards)
	max := "00"
	var tempCards []Card
	for _, cards := range possibleCardsList {
		temp := Convert5CardsToPatterns(cards)
		if temp > max {
			tempCards = cards
			max = temp
		}
	}
	
	return max, tempCards
}