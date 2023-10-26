// shell main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"AgusDOLARD/theo/builtin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "$HOME/.config/theo/theorc", "load the config file")
}

func main() {
	err := loadConfigFile(os.ExpandEnv(configPath))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		pwd := os.Getenv("PWD")
		directory := strings.Split(pwd, "/")
		fmt.Printf("%s > ", directory[len(directory)-1])
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		err = execCommand(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}

func execCommand(input string) error {
	// Remove the newline character.
	input = os.ExpandEnv(strings.TrimSuffix(input, "\n"))

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		return builtin.Cd(args[1:])
	case "addpath":
		return builtin.AddPath(args[1:])
	case "setenv":
		return builtin.SetEnviromentVariable(args[1:])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func loadConfigFile(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		execCommand(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
