package cli

import (
	"farming_game/pkg/state"
	"fmt"
)

func PrintStatus(state state.GameState) {
	fmt.Print("Your balance: ", state.Money,
	"\nDay: ", state.Day,
	"\n", state.Season,
	"\nWeather: ", state.Weather.Name)
}