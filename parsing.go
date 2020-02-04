package main

import (
	"regexp"

	"github.com/jessevdk/go-flags"
)

type programArgs struct {
	Date   string `short:"d" long:"date" description:"Date of your travel"`
	Number string `short:"n" long:"number" description:"Flight/Train number"`
	List   bool   `short:"l" long:"list" description:"List all watching records"`
	Add    bool   `short:"a" long:"add" description:"Add a travel record"`
	Delete bool   `short:"x" long:"delete" description:"Delete a travel record"`
}

func parseArgs() programArgs {
	var args programArgs

	_, err := flags.Parse(&args)
	if err != nil {
		panic(err)
	}

	return args
}

func parseInfo() TrafficInfo {
	args := parseArgs()

	i := TrafficInfo{
		Date:   parseDate(args.Date),
		Number: args.Number,
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
