package main

import (
	"fmt"
	"os"
	"os/exec"
)

type GoEnvAction struct {
	runFile string
}

func (act GoEnvAction) DoAction() error {
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Printf("Making go environment in %s...\n", currDir)
	if err := makeRunScript(act.runFile); err != nil {
		return err
	}
	if err := permitRunScript(act.runFile); err != nil {
		return err
	}

	return nil
}

func makeRunScript(runFile string) error {
	// Run file creation and writing
	f, err := os.Create("./" + runFile + ".sh")
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString("clear &&\ngo build -o ./a.out &&\n./a.out")

	return nil
}

func permitRunScript(runFile string) error {
	// Run file permissions
	cmd := exec.Command("chmod", "+x", "./"+runFile+".sh")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (act GoEnvAction) GetHelp() string {
	return "builds a go environment in the current directory"
}
