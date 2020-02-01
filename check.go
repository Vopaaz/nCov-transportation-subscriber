package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func check() {
	records := getData()
	if len(records) < 350 {
		panic("Number of records incorrect. The data API has probably changed.")
	}

	watching := readCurrentWatching()
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(watching); j++ {
			checkOne(records[i], watching[j])
		}
	}

	fmt.Println("Check successful.")
}

func checkOne(record requestRecord, watching TrafficInfo) {
	if len(record.Date) != len("2020-01-26") {
		panic("Format of date field in the data API has probably changed.")
	}

	if strings.Compare(record.Date, watching.Date) == 0 && strings.Compare(record.Number, watching.Number) == 0 {
		fmt.Println("[WARNING] " + record.Date + " " + record.Number + " IS REPORTED!")
	}
}

type requestSchema struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data []requestRecord `json:"data"`
}

type requestRecord struct {
	Date   string `json:"t_date"`
	Number string `json:"t_no"`
}

func getData() []requestRecord {
	url := "https://2019ncov.nosugartech.com/data.json"

	rs, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	bytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	var schema requestSchema
	json.Unmarshal(bytes, &schema)
	return schema.Data
}
