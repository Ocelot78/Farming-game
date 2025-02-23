package cli

import (
	"fmt"
	"os"
	
	"github.com/olekukonko/tablewriter"
	"strings"

	"farming_game/pkg/state"
)

type Command struct {
	Action 	string
	Target	string
}
func ParseCommand(input string) Command {
	Parser := strings.Split(input ,"-",)
	fmt.Print(input)
	fmt.Print(Parser)
	return Command{Action: Parser[0], Target: Parser[1]}
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
func ShopBuy() {
	
}
func ExitGame() {
	os.Exit(0)
}
func Disabled() {
	fmt.Print("This feature is disabled\n")
}