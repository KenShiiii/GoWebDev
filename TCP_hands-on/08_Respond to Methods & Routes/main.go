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
			// HEADER LINE
			m = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("***METHOD*** ", m)
			fmt.Println("***URI*** ", uri)
			switch {
			case m == "GET" && uri == "/":
				index(conn)
			case m == "GET" && uri == "/apply":
				apply(conn)
			case m == "POST" && uri == "/apply":
				applyPost(conn)
			default:
				index(conn)
			}
			i++
		}
		fmt.Println(ln)
	}

	defer conn.Close()
}

func index(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<a href="/"> <h1>Index</h1></a>
			<a href="/apply"> <h1>apply</h1></a>

		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	return
}

func apply(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<a href="/"><h1>Index</h1></a>
			<a href="/apply"><h1>apply</h1></a>
			<form action="/apply" method="POST">
  				<input type="submit" name="applyBtn" value="ClickToApply">
			</form>

		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	return

}

func applyPost(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<a href="/"> <h1>Index</h1></a>
			<a href="/apply"> <h1>apply</h1></a>
			<strong>apply Processed!</strong>

		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	return
}
