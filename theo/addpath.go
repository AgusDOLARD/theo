package theo

import (
	"fmt"
	"os"
)

// built in command to append paths to $PATH
func AddPath(args []string) error {
	if len(args) == 0 {
		return newError("addpath", "path/to/directory")
	}
	path := os.Getenv("PATH")
	return os.Setenv("PATH", fmt.Sprintf("%s:%s", path, args[0]))
}
