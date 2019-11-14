package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var config map[string]string

func Load() bool {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func Get(c string) string {
	return config[c]
}