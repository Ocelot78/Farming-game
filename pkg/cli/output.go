package cli

import (
	"fmt"

	"farming_game/pkg/state"
)

func PrintStatus(state state.GameState) {
	fmt.Print(
	"| Day: ", state.Day,
	" | ", state.Season,
	" | ", state.Weather.Name,
	" | Money: ", state.Money,
	"\n")
}	
func ReadCommand(state state.GameState) {
	var command string
	var target int
	fmt.Print("$")
	fmt.Scanln(&command)
	if command == "" {
		fmt.Print("You must enter a command \n")	
	}else {
		switch command {
		case "show":
			fmt.Print("What do you want to see?\n",
			"1.Day State\n",
			"2.Fields\n",
			"3.Machines\n",
			"4.Silo\n")
			fmt.Print("$")
			fmt.Scanln(&target)
			switch target {
			case 1:
				PrintStatus(state)
			case 2:
				ShowFields(state)
			case 3:
				Disabled()
			case 4:
				Disabled()
			default:
				fmt.Print("Unkown command try again.")
			}
		case "shop":
			fmt.Print("1.Buy machines\n2.Sell Machines")
			fmt.Print("$")
			fmt.Scanln(&target)
			switch target {
			case 1:
				
			case 2:
				Disabled()
			default:
				fmt.Print("Unkown command try again.")
			}
		case "exit":
			ExitGame()
		default:
			fmt.Print("Unkown command try again.")
		}
	}
}