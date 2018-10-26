package CoolProject
import ("fmt")


func PrintWinner(currentState *Current, maxSeatNumber []int) {
	for i := 0; i < len(maxSeatNumber); i++ {
		for j := 0; j < len(currentState.Players); j++ {

			if currentState.Players[j].SeatPosition == maxSeatNumber[i] {
				var sevenCardSlice []Card
				sevenCardSlice = append(sevenCardSlice, currentState.Players[j].Hands[0], currentState.Players[j].Hands[1])
				for m := 0; m < len(currentState.CommunityCard); m++ {
					sevenCardSlice = append(sevenCardSlice, currentState.CommunityCard[m])
				}
				_, tempCards := MaxPattern(sevenCardSlice)


				fmt.Print("The winner is ")
				fmt.Print(currentState.Players[j].Name)
				fmt.Print(" and the max pattern is ")
				fmt.Println(HandsToString(tempCards))
			}
		}

	}

}
