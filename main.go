package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var disp Dispatcher

func main() {
	args := os.Args[1:]
	disp := NewDispatcher()

	if len(args) == 0 {
		noArgs()
	}

	for _, command := range args {
		mapCommand(disp, command)
	}

	os.Exit(0)
}

func mapCommand(disp *Dispatcher, command string) {
	if command, err := disp.Map(command); err == nil {
		fmt.Printf("running %s command...\n", command.value)
		err := command.action.DoAction()

		if err != nil {
			log.Fatalf("Error executing %s\n%s", command.value, err)
		}

	} else {
		fmt.Printf("Error mapCommand: %s", err)
		//fmt.Printf("Invalid command: %s\n", command)
	}
}

func noArgs() {
	fmt.Println("No arguments supplied")
	printHelp()
	//fmt.Println(readInput())
}

func printHelp() {
	fmt.Println("Help:")

	for key, command := range getCommands() {
		fmt.Printf("%s%s: ", indent(1), key)
		fmt.Printf("%s\n\n", command.action.GetHelp())
	}
}

func indent(x int) string {
	indent := 3
	return strings.Repeat(" ", indent*x)
}

func readInput() string {
	fmt.Print("Input> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
