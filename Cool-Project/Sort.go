package CoolProject

/*
  Auther: Chengyang Nie
  Date: 10/11/2018 12:23
*/

/*
This function is used to sort five cards' num.
The input is a slice of 5 nums of 5 cards.
The output is a slice of sorted 5 nums.

*/

func Sort(numSlice []int) []int {
	for i := 0; i < len(numSlice) - 1; i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] < numSlice[j] {
				temp := numSlice[j]
				numSlice[j] = numSlice[i]
				numSlice[i] = temp
			}
		}
	}
	return numSlice
}