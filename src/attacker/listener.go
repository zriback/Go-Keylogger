package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 4444})
	check(err)
	defer conn.Close()

	var dataRead []byte
	for {
		// clear the console
		fmt.Println("\x1bc") // clears the terminal from the bottom

		buffer := make([]byte, 1024)
		n, addr, _ := conn.ReadFromUDP(buffer)
		dataRead = append(dataRead, buffer[:n]...) // extend

		fmt.Println(string(dataRead), "\tFrom", addr)
		time.Sleep(100 * time.Millisecond)
	}

}

// stop running if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}
