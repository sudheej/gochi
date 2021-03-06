package queue

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sudheej/gochi/plugins"
	"gopkg.in/yaml.v2"
)

//YamlParser would enable files to be parsed and the values to be used by the requester.
func YamlParser(fileName string) (result plugins.Data) {
	fmt.Println("Reading Configuration File..")
	str, err := ioutil.ReadFile(fileName)
	//var result plugins.Data
	yaml.Unmarshal([]byte(str), &result)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return result

}
