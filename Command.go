package main

import "github.com/Nastyyy/jenv/actions"

type Command struct {
	value  string
	action actions.Action
}

// TODO: Change later
func getCommands() map[string]*Command {
	commands := make(map[string]*Command)

	// Golang env
	commands["go"] = &Command{"go", actions.GoEnvAction{"run"}}

	// Python env
	commands["py"] = &Command{"py", actions.PythonEnvAction{}}

	return commands
}
