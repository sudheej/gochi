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
	fmt.Println("Parsing YAML file")
	str, err := ioutil.ReadFile(fileName)
	fmt.Printf("%+v\n", result)
	//var result plugins.Data
	yaml.Unmarshal([]byte(str), &result)
	fmt.Println(result)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return result

}
