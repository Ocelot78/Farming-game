package main

import (
	"farming_game/pkg/cli"
	//"farming_game/pkg/state"
	"farming_game/pkg/storage"
	//"fmt"
)

func main() {
	gameState := storage.LoadSave()
	cli.PrintStatus(gameState)
	for {
		cli.ReadCommand(gameState)
	}

}