package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)

	}
}

func serve(conn net.Conn) {
	i := 0
	m := ""
	uri := ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			fmt.Println("END OF HTTP HEADER")
			break
		}
		if i == 0 {
			m = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("***METHOD*** ", m)
			fmt.Println("***URI*** ", uri)
			i++
		}
		fmt.Println(ln)
	}
	body := "Method: " + m + "<br>URI: " + uri + "<br><br><h1><HOLY COW THIS IS LOW LEVEL</h1>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	defer conn.Close()
}

//This is answer from Todd

//body := `
//		<!DOCTYPE html>
//		<html lang="en">
//		<head>
//			<meta charset="UTF-8">
//			<title>Code Gangsta</title>
//		</head>
//		<body>
//			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
//		</body>
//		</html>
//	`
//io.WriteString(c, "HTTP/1.1 200 OK\r\n")
//fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
//fmt.Fprint(c, "Content-Type: text/html\r\n")
//io.WriteString(c, "\r\n")
//io.WriteString(c, body)