package cli

import (
	"fmt"
	"strings"
)

type Command struct {
	Action 	string
	Target	string
}

func ParseCommand(input string) Command {
	Parser := strings.Split(input ," ",)
	fmt.Print(input)
	fmt.Print(Parser)
	return Command{Action: Parser[0], Target: Parser[1]}
}