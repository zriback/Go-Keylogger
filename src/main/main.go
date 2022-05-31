package main

import (
	"fmt"
)

func main() {

	k := createKeylogger()

	fmt.Println(k)

	k.setSaveToFile("file path here")

	fmt.Println(k)

}

// create a keylogger with the default values
func createKeylogger() keylogger {
	return keylogger{}
}
