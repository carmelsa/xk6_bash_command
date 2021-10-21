package xk6_bash_command

import (
	"fmt"
	"go.k6.io/k6/js/modules"
	"os/exec"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/redis".
func init() {
	modules.Register("k6/x/bash_command", new(Bash))
}

type Bash struct{}

func (b *Bash) Exec(command string, args []string) string {
	cmd := exec.Command(command, args...)
	res, err := cmd.Output()
	if err != nil {
		commandString := fmt.Sprintf("%v %s", command, args)
		fmt.Println(fmt.Sprintf("error executing command %v : %v", commandString, err))
	}
	return string(res)
}
