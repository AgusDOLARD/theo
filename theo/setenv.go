package theo

import (
	"os"
)

// built in command to set an enviroment variable
func SetEnviromentVariable(args []string) error {
	if len(args) != 2 {
		return newError("setenv", "KEY value")
	}
	return os.Setenv(args[0], args[1])
}
