package CoolProject

/*this function print the welcome information for the this game
TEXAS Holdem @Yichen */
import (
	"bufio"
	"fmt"
	"os"
)

func WelcomeGame(currentBoard *Current) {
	fmt.Println("♣ ♦ ♥ ♠")
	fmt.Println("Welcome to Texas Holdem Game!!")
	fmt.Println("♣ ♦ ♥ ♠")
	fmt.Println("Are you ready to play the game?")
	RepeatInput("Enter")

	ClearScreen()
	for i, _ := range currentBoard.Players {
		fmt.Println("Your seat position is ", i+1, ". Please enter your nickname:")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = name[:len(name)-1]
		(*currentBoard).Players[i].Name = name
	}

}
