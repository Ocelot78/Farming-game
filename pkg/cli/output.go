package cli

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
	fmt.Scanln(&command)
	if command == "" {
		fmt.Print("You must enter a command")	
	}else {
		parsedCommand := ParseCommand(command)	
		switch parsedCommand.Action{
		case "show":
			switch parsedCommand.Target{
				case "fields":
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader([]string{"ID", "Crop", "Growth", "Size"})
					for _, Field := range state.Fields {
						row := []string{fmt.Sprint(Field.ID),Field.Crop, fmt.Sprint(Field.Growth), fmt.Sprint(Field.Size)}
						table.Append(row)
					}
					table.Render()
			}
		}
	}
}