package main

import (
	"net"
	"strconv"
)

const PORT = 50000

func main() {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if server == nil {
		panic("couldn't start listening: " + err.Error())
	}
	conns := handler.clientConns(server)
	for {
		go handler.handle(<-conns)
	}
}
