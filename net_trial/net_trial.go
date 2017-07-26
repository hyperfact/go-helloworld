package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

func hanleRead(con net.Conn, buf bytes.Buffer) {
	_, err := buf.ReadString('\n')
	if err != nil {
		return
	}

	fmt.Println()
}

func cmpIgnoreCase(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func handleConnection(con net.Conn) {
	defer func() {
		con.Close()
		log.Printf("connection closed: %v\n", con.RemoteAddr())
	}()

	input := bufio.NewScanner(con)
	for input.Scan() {
		t := input.Text()
		if cmpIgnoreCase(t, "exit") || cmpIgnoreCase(t, "quit") {
			break
		}
		fmt.Fprintln(con, "echo:", t)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal("Listen failed")
		return
	}

	defer ln.Close()

	for {
		con, err := ln.Accept()
		if err != nil {
			continue
		}

		log.Printf("new connection: %v\n", con.RemoteAddr())
		go handleConnection(con)
	}
}
