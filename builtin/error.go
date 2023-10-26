package builtin

import (
	"fmt"
)

func newError(command string, error string) error {
	return fmt.Errorf("%s: usage -> %s %s", command, command, error)
}
