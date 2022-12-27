package main

import (
	// scanner "github.com/gerguevara/go-net-chat/port"
	"flag"
	"fmt"
	"log"
	"net"

	chat "github.com/gerguevara/go-net-chat/chat"
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

func main() {
	// flag.Parse()
	// scanner.VerifyPortsCurrently()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go chat.Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go chat.HandleConnection(conn)
	}
}
