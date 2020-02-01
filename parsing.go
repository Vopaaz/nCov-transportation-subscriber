package main

import (
	"os"
	"regexp"

	"github.com/jessevdk/go-flags"
)

func isRegister() bool {
	return len(os.Args) != 1
}

func parseInfo() TrafficInfo {
	var opts struct {
		Date   string `short:"d" long:"date" description:"Date of your travel"`
		Number string `short:"n" long:"number" description:"Flight/Train number"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	i := TrafficInfo{
		Date:   parseDate(opts.Date),
		Number: opts.Number,
	}

	return i
}

func leftPad0(s string) string {
	if len(s) == 1 {
		s = "0" + s
	}
	if len(s) != 2 {
		panic("Date parsing error")
	}
	return s
}

func parseDate(date string) string {
	re := regexp.MustCompile("([0-9]{1,2})(?:-|/|\\\\)([0-9]{1,2})")
	res := re.FindAllStringSubmatch(date, -1)
	return "2020-" + leftPad0(res[0][1]) + "-" + leftPad0(res[0][2])
}
