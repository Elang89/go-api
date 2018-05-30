package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	MongoConnection struct {
		Host     string `json:"Host"`
		Database string `json:"Database"`
		User     string `json:"User"`
		Pwd      string `json:"Pwd"`
	} `json:"MongoConnection"`
}

var config Configuration

// NewConfiguration creates a configuration object with the configuration options found in config.json
func NewConfiguration() Configuration {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opening config.json...")

	fileData, _ := ioutil.ReadAll(file)

	fmt.Println("Reading config.json...")

	err = json.Unmarshal(fileData, &config)

	if err != nil {
		fmt.Println(err)
	}

	return config
}
