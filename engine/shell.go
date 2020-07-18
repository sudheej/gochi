package engine

import (
	"fmt"
	"os/exec"
)

//MavenRunner is the function is to run Maven commands
func MavenRunner() {
	executor("mvn", "-v")
}

func executor(cmdx string, param string) {
	app := cmdx
	arg0 := param

	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
