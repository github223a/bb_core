package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const appEnv = "APP_ENV"
const configPath = "./config.json"
const configTestPath = "./config.test.json"
const configDevelopmentPath = "./config.development.json"

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadConfig() interface{} {
	var config interface{}
	var err error
	var configFile *(os.File)

	if fileExists(configPath) == true {
		os.Setenv(appEnv, "production")
		configFile, err = os.Open(configPath)

	} else if fileExists(configDevelopmentPath) == true {
		os.Setenv(appEnv, "development")
		configFile, err = os.Open(configDevelopmentPath)

	} else {
		os.Setenv(appEnv, "test")
		configFile, err = os.Open(configTestPath)
	}

	FailOnError(HEADER_APPLICATION_MESSAGE, "Error on open config file", err)
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)
	uErr := json.Unmarshal([]byte(bytes), &config)
	FailOnError(HEADER_APPLICATION_MESSAGE, "Error on unmarhal config file.", uErr)

	return config
}
