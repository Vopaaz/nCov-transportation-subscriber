package main

import (
	"os"
)

func isCheck() bool {
	return len(os.Args) == 1
}

func operationType() string {

	args := parseArgs()

	if args.List && !args.Add && !args.Delete {
		return "l"
	} else if !args.List && args.Add && !args.Delete {
		return "a"
	} else if !args.List && !args.Add && args.Delete {
		return "d"
	} else if !args.List && !args.Add && !args.Delete {
		return "a"
	} else {
		panic("Choose only one from -l/--list, -a/--add, -x/--delete only.")
	}
}
