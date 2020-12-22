package main

import "github.com/Nastyyy/jmake/actions"

type Command struct {
	value  string
	action actions.Action
}

// TODO: Change later
func getCommands() map[string]Command {
	commands := make(map[string]Command)
	commands["go"] = Command{"go", actions.GoEnvAction{"jenv-go-run"}}
	commands["python"] = Command{"python", actions.PythonEnvAction{"jenv-py-run"}}

	return commands
}
