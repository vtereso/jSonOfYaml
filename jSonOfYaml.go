package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

var (
	o = flag.String("o", "stdout", "The file to output the json into.")
)

// Takes a single yaml file argument and pretty prints it as json
func main() {
	flag.Parse()
	os.Args = flag.Args()
	if len(os.Args) != 1 {
		fmt.Println("Filename was not passed")
		return
	}
	yamlFile := os.Args[0]
	yamlBytes, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	var out bytes.Buffer
	json.Indent(&out, jsonBytes, "", "	")
	if err != nil {
		fmt.Println(err)
		return
	}

	var outFile *os.File
	switch *o {
	case "stdout":
		outFile = os.Stdout
	default:
		outFile, err = os.OpenFile(*o, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Fprintf(outFile, out.String())
}
