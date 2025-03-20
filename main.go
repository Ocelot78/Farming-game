package main

import (
	"fmt"

	"farming_game/pkg/cli"
	//"farming_game/pkg/state"
	"farming_game/pkg/storage"
)

func main() {
	gameState := storage.LoadSave()
	shopTractors:= storage.LoadTractors()
	cli.PrintStatus(gameState)
	fmt.Print("Enter help for more info!\n")
	for {
		cli.ReadCommand(gameState, shopTractors)
	}

}