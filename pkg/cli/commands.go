package cli

import (
	"fmt"
	"os"
	
	"github.com/olekukonko/tablewriter"


	"farming_game/pkg/state"
)

type Command struct {
	Action 	string
	Target	string
}
func ShowFields(state state.GameState) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Crop", "Growth", "Size"})
	for _, Field := range state.Fields {
		row := []string{fmt.Sprint(Field.ID),Field.Crop, fmt.Sprint(Field.Growth), fmt.Sprint(Field.Size)}
		table.Append(row)
	}
	table.Render()
}
func Tractors(shop state.Shop) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID","Name", "Power index", "Price"})
	for _, Tractor := range shop.Tractors {
		row := []string{fmt.Sprint(Tractor.ID),fmt.Sprint(Tractor.Name),fmt.Sprint(Tractor.PowerIndex),fmt.Sprint(Tractor.Price)}
		table.Append(row)
	}
	table.Render()
}
func ExitGame() {
	os.Exit(0)
}
func Disabled() {
	fmt.Print("This feature is disabled\n")
}