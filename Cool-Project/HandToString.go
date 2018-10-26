package CoolProject

import (
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
