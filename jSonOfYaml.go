package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

// Takes a single yaml file argument and pretty prints it as json
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Filename was not passed")
		return
	}
	file := os.Args[1]
	yamlBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		fmt.Println("Error converting yaml to json:", err)
		return
	}
	var out bytes.Buffer
	json.Indent(&out, jsonBytes, "", "	")
	if err != nil {
		fmt.Println("Error indenting json:", err)
		return
	}
	fmt.Println(out.String())
}
