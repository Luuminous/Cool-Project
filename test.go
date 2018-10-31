package main

import (
	"CoolProject"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numPlayers, _ := strconv.Atoi(os.Args[1])
	initialChips, _ := strconv.Atoi(os.Args[2])
	numTurns, _ := strconv.Atoi(os.Args[3])
	/*
	numPlayers := 4
	initialChips := 1000
	numTurns := 5
	*/
	CoolProject.PlayGame(numPlayers, initialChips, numTurns)
}
