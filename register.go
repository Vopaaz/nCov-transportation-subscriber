package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func register() {

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, os.ModeDir)
	}

	var watch []TrafficInfo

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		watch = make([]TrafficInfo, 0)
	} else {
		watch = readCurrentWatching()
	}

	watch = append(watch, parseInfo())

	file, _ := json.MarshalIndent(jsonSchema{TrafficInfoList: watch}, "", " ")

	err := ioutil.WriteFile(configFile, file, os.FileMode(0644))

	if err != nil {
		panic(err)
	}
}
