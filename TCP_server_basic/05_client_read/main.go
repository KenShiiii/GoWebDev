package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	xb, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(xb))
}
