package main

import (
	"flag"
	"fmt"
	"vim-syntax/cursor"
	"vim-syntax/global"
	"vim-syntax/help"
	"vim-syntax/insert"
)

type CommandsForLife interface {
	GetCommands() map[string]string
}

func main() {
	command := flag.String("c", "", "Command to execute")
	flag.StringVar(command, "command", "", "Command to execute")
	flag.Parse()

	var c CommandsForLife

	switch *command {
	case global.GlobalName:
		c = global.NewGlobal()
	case cursor.CursorName:
		c = cursor.NewCursor()
	case insert.InsertName:
		c = insert.NewInsert()
	default:
		c = help.NewHelp()
	}

	commands := c.GetCommands()
	for k, command := range commands {
		fmt.Println(fmt.Sprintf("With %s", k), fmt.Sprintf("-- you can %s", command))
	}
}
