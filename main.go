//Chat is a simple synchronized program to chat between two people,
// one being the host and the other the guest.
package main

import (
	"chat/server"
	"flag"
	"os"
)

func main() {

	var isHost bool //either a host or guest

	flag.BoolVar(&isHost, "listen", false, "Listen on the specified ip address")
	flag.Parse()

	if isHost {
		//go run main -listen <ip-address>
		connIP := os.Args[2]
		server.RunHost(connIP)
	} else {
		//go run main.go <ip-address>
		connIP := os.Args[1]
		server.RunGuest(connIP)
	}
}
