package main

type Command struct {
	value  string
	action Action
}

// TODO: Change later
func getCommands() map[string]Command {
	commands := make(map[string]Command)
	command := Command{"go", GoEnvAction{"test-run"}}

	commands[command.value] = command

	return commands
}

type Action interface {
	DoAction() error
	GetHelp() string
}
