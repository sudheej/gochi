package engine

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//YamlParser would enable files to be parsed and the values to be used by the requester.
func YamlParser(fileName string) {
	fmt.Println("Parsing YAML file")

	data, err := ioutil.ReadFile(fileName)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	for k, v := range m {
		fmt.Println("k:", k, "v:", v)
	}

	/*
		d, err := yaml.Marshal(&m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m dump:\n%s\n\n", string(d))
	*/
}
