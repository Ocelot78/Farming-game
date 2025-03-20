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
func ShowTractors(state state.GameState){
	fmt.Print(state.OwnedT)
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
func BuyTractors(state state.GameState, shop state.Shop, userInput int){
	for _, Tractor := range shop.Tractors {
		if Tractor.ID == userInput && Tractor.Price <= state.Money{
			state.OwnedT = Tractor
			break
		}
	}
}
func ExitGame() {
	os.Exit(0)
}
func Unknown() {
	fmt.Print("Unknown command try again.\n")
}
func Disabled() {
	fmt.Print("This feature is disabled\n")
}