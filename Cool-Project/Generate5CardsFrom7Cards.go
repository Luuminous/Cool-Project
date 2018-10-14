package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/11/2018
*/

/*
This function is used to genenrate 5 cards from 7 cards
The input is a slice of card.
The output is a two D slice of card.
I will use brute force method to list all the combinations.
*/

func Generate5CardsFrom7Cards(cardSlice []Card) [][]Card {
	var slice [][]int
	for i := 0; i < len(cardSlice) - 1; i++ {
		for j := i + 1; j < len(cardSlice); j++ {
			temp := [] int {i, j} 
			slice = append(slice, temp)
		}
	}
    
    var newSlice [][]Card
    for i := 0; i < len(slice); i++ {
    	var tempSlice []Card
    	for j := 0; j < len(cardSlice); j++ {
    		if (j != slice[i][0]) && (j != slice[i][1]) {
    			tempSlice = append(tempSlice, cardSlice[j])
    		}
    	}
    	newSlice = append(newSlice, tempSlice)
    }
    return newSlice

}