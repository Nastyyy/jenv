package main

import "github.com/Nastyyy/jmake/actions"

type Command struct {
	value  string
	action actions.Action
}

// TODO: Change later
func getCommands() map[string]Command {
	commands := make(map[string]Command)
	command := Command{"go", actions.GoEnvAction{"test-run"}}

	commands[command.value] = command

	return commands
}
