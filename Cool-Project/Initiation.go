/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

/*
This function is to create a new 52 cards in a slice.
*/
func Initiation() []Card {
	var ans []Card
	for i := 2; i <= 14; i++ {
		for j := 0; j <= 3; j++ {
			var temp Card
			temp.Num = i
			temp.Color = j
			ans = append(ans, temp)
		}
	}
	return ans
}