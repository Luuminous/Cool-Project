/*
Author: Yichen Mo
Date: Oct. 31st, 2018
 */

package CoolProject

func IsInString(s string, array []string) bool {
	for _, val := range array {
		if val == s {
			return true
		}
	}
	return false
}
