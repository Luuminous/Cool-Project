

package CoolProject

/*
  Author: Chengyang Nie
  Date: 10/11/2018
*/

/*
This function is used to check if the 5 nums is a straight.
The input of the function is the sorted five nums in a slice.
The output of the funciton is a bool.

Check, if the distance of two nums is 1, if not return false.
*/
func IsStraight(sortedSlice []int) bool {
	//check if this is bicycle
	if sortedSlice == nil {
		panic("The input sortedSlice is empty")
	}
	if sortedSlice[0] == 14 && sortedSlice[1] == 5 && sortedSlice[2] == 4 && sortedSlice[3] == 3 && sortedSlice[4] == 2{
		return true
	}
	for i := 0; i < len(sortedSlice) - 1; i++ {
		//check the distance between two is 1
		if sortedSlice[i] - sortedSlice[i + 1] == 1 {
			continue
		} else {
			return false
		}
	}
	return true
}