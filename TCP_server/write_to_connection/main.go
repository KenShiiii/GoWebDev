package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for  {
		conn, err :=  li.Accept()
		if err != nil{
			log.Println(err)
			continue
		}

		fmt.Fprintln(conn, "\nHello from TCP server\n\n")
		io.WriteString(conn, "Hope you are well")

		conn.Close()
	}
}
