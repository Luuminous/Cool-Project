package main

import (
	"CoolProject"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cardList := make([]CoolProject.Card, 2)
	cardList[0].Num = 11
	cardList[0].Color = 1
	cardList[1].Num = 10
	cardList[1].Color = 0
	CoolProject.CalculateProbability(100000)
}
