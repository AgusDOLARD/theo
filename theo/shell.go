package theo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	configFile      string
	builtinCommands []*Command
}

type CommandExecute func([]string) error

type Command struct {
	Name string
	Func CommandExecute
}

func (c *Command) Run(args []string) error {
	return c.Func(args)
}

func New(configFile string) *Shell {
	return &Shell{configFile: configFile}
}

func (s *Shell) BuiltInCommand(name string, f CommandExecute) {
	s.builtinCommands = append(s.builtinCommands, &Command{Name: name, Func: f})
}

func (s *Shell) Init() {
	err := s.loadConfigFile(s.configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		prompt := "> "
		fmt.Print(prompt)
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = s.execCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func (s *Shell) execCommand(input string) error {
	input = os.ExpandEnv(strings.TrimSuffix(input, "\n"))

	args := strings.Split(input, " ")

	if args[0] == "exit" {
		os.Exit(1)
	}

	cmd := exec.Command(args[0], args[1:]...)

	for _, command := range s.builtinCommands {
		if command.Name == args[0] {
			return command.Run(args[1:])
		}
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (s *Shell) loadConfigFile(fp string) error {
	f, err := os.Open(os.ExpandEnv(fp))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		err := s.execCommand(scanner.Text())
		if err != nil {
			return err
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
