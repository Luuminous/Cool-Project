/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

/*
This function is to delete the out cards from the pool cards.
*/
func DeleteFromPool(pool []Card, out []Card) []Card {
	var ans []Card
	for i := 0; i < len(pool); i++ {
		check := true
		for j := 0; j < len(out); j++ {
			if (pool[i].Num == out[j].Num) && (pool[i].Color == out[j].Color) {
				check = false
			}
		}
		if check {
			ans = append(ans, pool[i])
		}
	}
	return ans
}
