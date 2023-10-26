package builtin

import (
	"os"
)

// built in command to change directory
func Cd(args []string) error {
	var path string
	if len(args) == 0 {
		path = os.Getenv("HOME")
	} else {
		path = args[1]
	}
	return os.Chdir(path)
}
