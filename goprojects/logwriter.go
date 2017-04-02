package main

import (
	"net"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", "hawk.l3.com:1902")
	if err != nil {
		panic("Failed to connect to hawk.l3.com:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")

}
