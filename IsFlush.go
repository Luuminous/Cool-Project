package coolProject

/*
  Author: Chengyang Nie
  Date: 10/11/2018 1:02
 */


/*
This function is used to check if these five cards are the same color
The input is a slice with 5 numbers which represents the color
The output is a bool

*/
func IsFlush(colorSlice []int) bool {
	for i := 0; i < len(colorSlice); i++ {
		if colorSlice[0] != colorSlice[i] {
			return false
		}
	}
	return true
}