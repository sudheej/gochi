package queue

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sudheej/gochi/engine"
	"gopkg.in/yaml.v2"
)

type data struct {
	Steps struct {
		Execute []execute `yaml:"execute"`
	} `yaml:"steps"`
}

type execute struct {
	MVN        string `yaml:"mvn"`
	Goals      string `yaml:"goals"`
	Concurrent bool   `yaml:"concurrent"`
}

//YamlParser would enable files to be parsed and the values to be used by the requester.
func YamlParser(fileName string) {
	fmt.Println("Parsing YAML file")
	str, err := ioutil.ReadFile(fileName)

	var result data
	yaml.Unmarshal([]byte(str), &result)
	fmt.Println(result)
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

	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
