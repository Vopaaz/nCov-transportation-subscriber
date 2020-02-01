package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

var appName string = "nCov-transportation-subscriber"
var home, _ = os.UserConfigDir()
var configDir = path.Join(home, appName)
var configFile = path.Join(configDir, "config.json")

// TrafficInfo is the information for train or flight
type TrafficInfo struct {
	Date   string `json:"date"`
	Number string `json:"number"`
}

type jsonSchema struct {
	TrafficInfoList []TrafficInfo `json:"info"`
}

func readCurrentWatching() []TrafficInfo {
	jsonFile, err := os.Open(configFile)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(jsonFile)
	var schema jsonSchema
	json.Unmarshal(bytes, &schema)
	return schema.TrafficInfoList
}
