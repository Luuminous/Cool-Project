package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/11/2018
*/

/*
This function is used to create random card from the pool
The input is the number of cards and a slice of Card pool
The output is a slice of cards.

use rand package shuffle function, then choose the first n elements.

*/

func CreateRandomCards(num int, cardPool []Card) []Card {
	rand.Shuffle(len(cardPool), func(i, j int) {
		cardPool[i], cardPool[j] = cardPool[j], cardPool[i]
	})
	var result []Card
	for i := 0; i < num; i++ {
		result = append(result, cardPool[i])
	}
	return result
}