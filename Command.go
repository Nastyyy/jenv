package main

import "github.com/Nastyyy/jenv/actions"

type Command struct {
	value  string
	action actions.Action
}

// TODO: Change later
func getCommands() map[string]*Command {
	commands := make(map[string]*Command)

	commands["go"] = &Command{"go", actions.GoEnvAction{"jenv-go-run"}}

	// Python
	pyCommand := &Command{"py", actions.PythonEnvAction{"jenv-py-run"}}
	commands["python"] = pyCommand
	commands["py"] = pyCommand

	return commands
}
