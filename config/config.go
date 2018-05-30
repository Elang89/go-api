package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration object with all settings
type Configuration struct {
	MongoConnection struct {
		Host     string `json:"Host"`
		Database string `json:"Database"`
		User     string `json:"User"`
		Pwd      string `json:"Pwd"`
	} `json:"MongoConnection"`
}

var config Configuration

// MongoHostName is a constant for the host name of the mongo service
const MongoHostName string = "mongodb://"

// NewConfiguration creates a configuration object with the configuration options found in config.json
func NewConfiguration() Configuration {
	file, err := os.Open("config.json")

	if err != nil {
		panic(err)
	}
	fmt.Println("Opening config.json...")

	fileData, _ := ioutil.ReadAll(file)

	fmt.Println("Reading config.json...")

	err = json.Unmarshal(fileData, &config)

	if err != nil {
		panic(err)
	}

	return config
}

// GenerateConnectionString returns a s string with all connection details for a MongoDB database
func GenerateConnectionString(c Configuration) string {
	connectionString := MongoHostName + c.MongoConnection.User + ":" + c.MongoConnection.Pwd + "@" + c.MongoConnection.Host
	return connectionString
}
