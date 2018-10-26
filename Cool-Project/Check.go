/* this is the code for five actions for Texas Holdem GameBoard
including Check()
@Yichen Mo 0ct.16.2018
*/
package CoolProject

/*Check is the action that take the Player and update the attributes in the Player also update the Current
 */
func Check(currentBoard *Current, player *Player) {
	(*player).OK = true
	(*currentBoard).PreEventsList = append((*currentBoard).PreEventsList, (player.Name + " chose to check"))
}
