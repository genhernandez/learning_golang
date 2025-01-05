package main

import "fmt"

type Command interface {
	Execute() string
}

type PrintCommand struct{}

func (p PrintCommand) Execute() string {
	return "Printing!"
}

type SaveCommand struct{}

func (s SaveCommand) Execute() string {
	return "Saving!"
}

func runCommands(commands []interface{}) {
	for _, cmd := range commands {
		if cmd, ok := cmd.(Command); ok {
			fmt.Println(cmd.Execute())
		}
	}
}

func main() {
	commands := []interface{}{
		PrintCommand{},
		"Not a command",
		2.1,
		SaveCommand{},
	}
	runCommands(commands)

}
