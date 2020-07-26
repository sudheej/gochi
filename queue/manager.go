package queue

import (
	"fmt"

	"github.com/sudheej/gochi/engine"
	"github.com/sudheej/gochi/plugins"
)

//Process would be the invocation point to read the configuration file and execution operations
func Process() {

	result := YamlParser(".gochi.yml")
	//Execution VCS operations

	for _, e := range result.Vcs.Git {
		fmt.Println("Cloning..", e.URL)
		plugins.GitClone(e)

	}

	for _, e := range result.Steps.Execute {
		fmt.Println()
		switch os := e.Concurrent; os {
		case false:
			fmt.Println("Serial Execution")
			engine.MavenRunner(e.Goals)
		case true:
			fmt.Println("Concurrent Execution Activated..")
		default:
			fmt.Println("Serial Execution")
		}

	}

}
