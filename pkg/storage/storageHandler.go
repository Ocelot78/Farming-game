package storage

import (
	"os"
	"encoding/json"
	"log"

	"farming_game/pkg/state"
)

func LoadSave() state.GameState {
	data, err := os.ReadFile("saves/save.json")
	if err != nil {
		log.Print("ERROR: FAILED TO READ SAVE GAME", err)
		return state.GameState{}
	}
	var gameState state.GameState
	err = json.Unmarshal(data, &gameState)
	if err != nil {
		log.Print("ERROR: FAILED TO UNMARSHAL SAVE GAME", err)
		return state.GameState{}
	}
	return gameState
}