package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("listen 8889...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	processor := Processor{Conn: conn}
	err := processor.process()
	if err != nil {
		fmt.Println(err)
	}
}
