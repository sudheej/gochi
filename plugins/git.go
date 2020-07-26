package plugins

import (
	"fmt"
	"os"

	"github.com/go-git/go-git"
)

//Git structure for mvn temporarily
type Git struct {
	URL string `yaml:"url"`
}

//GitClone will download source code
func GitClone(g Git) {
	// Clone the given repository to the given directory

	_, err := git.PlainClone("./.temp", false, &git.CloneOptions{
		URL:      g.URL,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
	}

}
