package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func manageWatchList() {
	switch operationType() {
	case "l":
		list()
	case "a":
		add()
	case "d":
		delete()
	}
}

func list() {
	watch := readCurrentWatching()
	fmt.Println("Current watching:")
	for i := 0; i < len(watch); i++ {
		fmt.Println(watch[i].Date + " " + watch[i].Number)
	}
}

func add() {
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
	mustWriteWatchListToFile(watch)
}

func mustWriteWatchListToFile(watch []TrafficInfo) {
	file, _ := json.MarshalIndent(jsonSchema{TrafficInfoList: watch}, "", " ")

	err := ioutil.WriteFile(configFile, file, os.FileMode(0644))

	if err != nil {
		panic(err)
	}
}

func delete() {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		panic("Record file does not exist. You are not watching any travels.")
	}

	newWatch := make([]TrafficInfo, 0)
	watch := readCurrentWatching()
	toDelete := parseInfo()

	for i := 0; i < len(watch); i++ {
		if watch[i].Date != toDelete.Date || watch[i].Number != toDelete.Number {
			newWatch = append(newWatch, watch[i])
		}
	}

	mustWriteWatchListToFile(newWatch)
}
