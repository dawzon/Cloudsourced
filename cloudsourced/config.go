package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const configFileName = "config.json"

type Config struct {
	MongoHost string
	MongoUser string
	MongoPass string
}

var config Config

func loadConfig() {

	jsonText, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonText, &config)
	if err != nil {
		log.Fatal(err)
	}
}

// func saveConfig() {

// 	err := ioutil.WriteFile(configFileName, []byte("This is a test"), 0600)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
