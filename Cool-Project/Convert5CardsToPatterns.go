/*
Author: Luuminous
Date: Oct. 11th, 2018
*/

package CoolProject

import (
	"strconv"
)

type Card struct {
	Num   int
	Color int
}

/*
This function is to convert a given integer to a string.
Add zero to one digital innteger.
*/
func ConvertIntToStr(x int) string {
	ans := strconv.Itoa(x)
	if len(ans) == 1 {
		ans = "0" + ans
	}
	return ans
}

/*
This function is to judge the pattern of the given five cards.
Input is a slice of card, output is a string, which can be compared by alphabet order.
*/
func Convert5CardsToPatterns(hands []Card) string {
	var numList []int // Record the num of five card.
	var colorList []int // Record the color of five card.
	for _, handCard := range hands {
		numList = append(numList, handCard.Num)
		colorList = append(colorList, handCard.Color)
	}
	sortedNumList := Sort(numList)
	if IsStraight(sortedNumList) {
		if IsFlush(colorList) {
			if (sortedNumList[0] == 14) && (sortedNumList[1] == 13) {
				// Royal flush
				ans := "9"
				return ans
			} else if (sortedNumList[0] == 14) {
				// Straight flush bicycle
				ans := "805"
				return ans
			} else {
				// Staight flush
				ans := "8"
				ans += ConvertIntToStr(sortedNumList[0])
				return ans
			}
		} else {
			if (sortedNumList[0] == 14) && (sortedNumList[1] == 13) {
				// Straight
				ans := "414"
				return ans
			} else if (sortedNumList[0] == 14) {
				// Straight bicycle
				ans := "405"
				return ans
			} else {
				// Straight
				ans := "4"
				ans += ConvertIntToStr(sortedNumList[0])
				return ans
			}
		}
	} else {
		duplicate, numDuplicate := Duplicate(sortedNumList)
		if duplicate[0] == 1 {
			if IsFlush(colorList) {
				// Flush
				ans := "5"
				for _, value := range sortedNumList {
					ans += ConvertIntToStr(value)
				}
				return ans
			} else {
				// High card
				ans := "0"
				for _, value := range sortedNumList {
					ans += ConvertIntToStr(value)
				}
				return ans
			}
		} else if duplicate[0] == 4 {
			// Four of a kind
			ans := "7"
			for _, value := range numDuplicate {
				ans += ConvertIntToStr(value)
			}
			return ans
		} else if duplicate[0] == 3 {
			if duplicate[1] == 2 {
				// Full house
				ans := "6"
				for _, value := range numDuplicate {
					ans += ConvertIntToStr(value)
				}
				return ans
			} else {
				// Three of a kind
				ans := "3"
				for _, value := range numDuplicate {
					ans += ConvertIntToStr(value)
				}
				return ans
			}
		} else if duplicate[0] == 2 {
			if duplicate[1] == 2 {
				// Two pair
				ans := "2"
				for _, value := range numDuplicate {
					ans += ConvertIntToStr(value)
				}
				return ans
			} else {
				// One pair
				ans := "1"
				for _, value := range numDuplicate {
					ans += ConvertIntToStr(value)
				}
				return ans
			}
		}
	}
	return "00" // Illegal default output
}