package plugins

import (
	"os"

	"github.com/sudheej/gochi/stdout"
	"gopkg.in/src-d/go-git.v4"
)

//Git structure for mvn temporarily
type Git struct {
	URL string `yaml:"url"`
}

//GitClone will download source code
func GitClone(g Git) {
	// Clone the given repository to the given directory

	stdout.ConsoleOut(true, "Git", "Cloning")
	_, err := git.PlainClone("./.temp", false, &git.CloneOptions{
		URL:      g.URL,
		Progress: os.Stdout,
	})

	if err != nil {
		//stdout.ConsoleError(false, "Git", err.Error())
		stdout.ConsoleUI(err.Error())
	} else {
		//stdout.ConsoleOut(false, "Git", "Done..!!")
		stdout.ConsoleUI("Download completed..")
	}

}
