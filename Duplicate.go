package CoolProject

/*Duplicate check whether there is the duplicate in the card
input is the slice of int, return two slices the value and its occurences
@Yichen Oct.11.2018
*/
func Duplicate(cards []int) ([]int, []int) {
	valueSlice := []int{}
	occurSlice := []int{}
	occur := make(map[int]int) //this is the map that key is the card number and value is the occuurences
	for i := 0; i < len(cards); i++ {
		if _, ok := occur[cards[i]]; ok {
			occur[cards[i]]++
		} else {
			occur[cards[i]] = 1
		}
	}
	maxNum := 0           // the max card value
	maxOccur := 0         // the occurences of Value
	for len(occur) != 0 { //the map is not empty
		maxNum, maxOccur = 0, 0
		for key, times := range occur {
			if times > maxOccur {
				maxOccur, maxNum = times, key
			} else if times == maxOccur && key > maxNum {
				maxNum, maxOccur = key, times
			}
		}
		valueSlice = append(valueSlice, maxNum)
		occurSlice = append(occurSlice, maxOccur)
		delete(occur, maxNum)
	}
	return occurSlice, valueSlice
}
