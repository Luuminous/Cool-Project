package CoolProject

import (
    "fmt"
    "bufio"
    "os"
)


/*RepeatInput promote the user to enter word to continue the next step */
func RepeatInput(str string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type " + str)
	action, _, _ := reader.ReadLine()
	aStr := string(action)
	for aStr != str {
		fmt.Println("Type " + str)
		action, _, _ = reader.ReadLine()
		aStr = string(action)
	}
}
