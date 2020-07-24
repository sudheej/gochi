package plugins

import (
	"fmt"
	"os"

	"github.com/go-git/go-git"
)

//GitClone will download source code
func GitClone() {
	// Clone the given repository to the given directory

	_, err := git.PlainClone("./download", false, &git.CloneOptions{
		URL:      "https://github.com/eiselems/spring-boot-microservices",
		Progress: os.Stdout,
	})

	if err == nil {
		fmt.Println(err)
	}

}
