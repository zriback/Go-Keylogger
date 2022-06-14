package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 4444})
	check(err)
	defer conn.Close()

	for {
		buffer := make([]byte, 512)
		n, addr, _ := conn.ReadFromUDP(buffer) // this pauses execution?
		_ = addr
		fmt.Print(string(buffer[:n]))
	}

}

// stop running if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}
