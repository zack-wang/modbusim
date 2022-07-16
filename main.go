package main

import (
	"flag"
	"log"
	"time"

	"github.com/tbrandon/mbserver"
)

var (
	host = flag.String("host", "127.0.0.1:1502", "modbus server address")
)

func main() {
	flag.Parse()
	serv := mbserver.NewServer()
	err := serv.ListenTCP(*host)
	if err != nil {
		log.Printf("%v\n", err)
		return
	} else {
		log.Println("Server ", *host, ` listening...`)
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
