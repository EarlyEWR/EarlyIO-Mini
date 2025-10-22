package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	log.Println("TCP echo server listening on :9000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "quit" {
			fmt.Fprintln(conn, "bye")
			return
		}
		fmt.Fprintf(conn, "echo: %s\n", line)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("conn scan error: %v", err)
	}
}
