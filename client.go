package main

import (
	"log"
	"net"
	"runtime"
)

func client(blockProgram chan bool) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ":8080")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("client: connection error: %v", err)
	}
	for i := 0; i < 5; i++ {
		conn.Write([]byte("HELLO maillennia"))
	}
	blockProgram <- true
	conn.Close()
}

func main() {
	runtime.GOMAXPROCS(2)

	conns := 1
	blockProgram := make(chan bool, conns)
	for i := 0; i < conns; i++ {
		go client(blockProgram)
	}
	for i := 0; i < int(conns); i++ {
		<-blockProgram
	}
}
