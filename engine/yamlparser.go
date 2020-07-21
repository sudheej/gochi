package engine

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func walk(v interface{}) {
	switch v := v.(type) {
	case []interface{}:
		for i, v := range v {
			fmt.Println("index:", i)
			walk(v)
		}
	case map[interface{}]interface{}:
		for k, v := range v {
			fmt.Println("key:", k)
			walk(v)
		}
	default:
		fmt.Println(v)
	}
}

//YamlParser would enable files to be parsed and the values to be used by the requester.
func YamlParser(fileName string) {
	fmt.Println("Parsing YAML file")

	data, err := ioutil.ReadFile(fileName)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	walk(m)
}
