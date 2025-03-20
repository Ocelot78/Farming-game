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
func ReadCommand(state state.GameState, shop state.Shop) {
	var command string
	var target int
	var targetString string
	fmt.Print("farmer@farm-os:~$")
	fmt.Scanln(&command)
	if command == "" {
		fmt.Print("You must enter a command \n")	
	}else {
		switch command {
		case "help": 
			fmt.Print("HELP\n",
			"help - enter help menu\n",
			"show - enter show menu\n",
			"shop - enter shop menu\n",
			"exit - exit game\n",
		)
		case "show":
			fmt.Print("What do you want to see?\n",
			"1.Day State\n",
			"2.Fields\n",
			"3.Machines\n",
			"4.Silo\n",
			"0 to go back\n")
			fmt.Print("farmer@farm-os:~/show$")
			fmt.Scanln(&target)
			switch target {
			case 1:
				PrintStatus(state)
			case 2:
				ShowFields(state)
			case 3:
				ShowTractors(state)
			case 4:
				Disabled()
			case 0:
				break
			default:
				Unknown()
			}
		case "shop":
			var machineID int
			fmt.Print("1.Buy machines\n2.Sell Machines\n")
			fmt.Print("farmer@farm-os:~/shop$")
			fmt.Scanln(&target)
			switch target {
			case 1:
				fmt.Print("What do you want to buy?\n",
				"1.Tractors\n",
				"2.Harvesters\n")
				fmt.Print("farmer@farm-os:~/shop/buy$")
				fmt.Scanln(&target)
				switch target {
				case 1:
					fmt.Print("Select thing you want to buy\n",
					"0 - Exit shop\n")
					Tractors(shop)
					fmt.Print("farmer@farm-os:~/shop/buy/tractors$")
					fmt.Scanln(&targetString)
					switch targetString {
					case "0":
						break
					case "buy":
						fmt.Print("Enter ID of a tractor you want to buy.\n")
						fmt.Print("farmer@farm-os:~/shop/buy/tractors$")
						fmt.Scanln(&machineID)
						BuyTractors(state, shop, machineID )
					}
				}
			case 2:
				Disabled()
			default:
				Unknown()
			}
		case "exit":
			ExitGame()
		default:
			Unknown()
		}
	}
}