/*
Author: Yichen Mo
Date: Oct. 3ist, 2018
 */

package CoolProject

import "fmt"

func ClearScreen() {
	//clean the screen
	for i := 0; i < 56; i++ {
		fmt.Println()
	}
}
