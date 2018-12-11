package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func ReadConfig() CoreConfig {
	var config CoreConfig
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

	FailOnError(err, "Error on open config file")
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)
	uErr := json.Unmarshal([]byte(bytes), &config)

	FailOnError(uErr, "Error on unmarhal config file.")
	return config
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s", msg)
		panic(err)
	}
}
