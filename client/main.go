package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Client Start")

	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	username := get_input("Username: ")
	fmt.Fprintf(conn, username + "\n")

	go get_server_messages(&conn, &username)
	for {
		message := get_input(username + " > ")
		fmt.Fprintf(conn, message + "\n")
	}
}

func get_server_messages(conn *net.Conn, username *string) {
	for {
		message, _ := bufio.NewReader(*conn).ReadString('\n')
		fmt.Print("\n" + message + *username + " > ")
	}
}

func get_input(question string) string {
	reader := bufio.NewReader(os.Stdin)
  fmt.Print(question)
	input, _ := reader.ReadString('\n')
	return strings.TrimRight(input, "\r\n")
}
