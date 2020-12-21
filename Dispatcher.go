package main

import "fmt"

type Dispatcher struct {
	Commands map[string]Command
	current  *Command
	previous *Command
}

func (d *Dispatcher) Map(command string) (*Command, error) {
	if dispCommand, ok := d.Commands[command]; ok {
		return &dispCommand, nil
	}

	return nil, fmt.Errorf("Invalid index on command map: %s", command)
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{Commands: getCommands()}
}
