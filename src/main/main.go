package main

import (
	"fmt"
)

func main() {

	k := createKeylogger()
	fmt.Println("Start logging:")
	k.startLogging(20)
	fmt.Println("DONE")
	//fmt.Println(k.log)

	// prints out key presses in order
	for _, key := range k.log {
		if key.down {
			fmt.Print(" ", key.name, " ")
		}
	}

}

// create a keylogger with the default values
func createKeylogger() keylogger {
	return keylogger{}
}
