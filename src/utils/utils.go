package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func ParseJSONFile(path string, data interface{}) {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read file")
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonFile, data)
	if err != nil {
		fmt.Println("Failed to unmarshall file")
		log.Fatal(err)
	}
}
