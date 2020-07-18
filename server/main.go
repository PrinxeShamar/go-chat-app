package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)


func main() {
	fmt.Println("Starting Server")
	// Creates The Server
	ln, _ := net.Listen("tcp", ":8000")
	// Map that stores different connections
	connections := make(map[net.Conn]string)
	for {
		connection, _ := ln.Accept()
		go new_connection(&connections, &connection)
	}
}

func new_connection(connections_addr *map[net.Conn]string, connection_addr *net.Conn) {
	connections := *connections_addr
	connection := *connection_addr
	username, _ := get_message(&connection)
	connections[connection] = username
	fmt.Println("IP:", connection.RemoteAddr().String(), "\nUsername:", username)
	send_message(&connections, &connection, username + " joined the chat!")
	for {
		message, err := get_message(&connection)
		if err != nil {
			fmt.Println("IP:", connection.RemoteAddr().String(), "Closed")
			connection.Close()
			delete(connections, connection)
			send_message(&connections, &connection, username + " left the chat!")
			return 
		}
		send_message(&connections, &connection, username + " > " + message)
	}
}

func send_message(connections_addr *map[net.Conn]string, connection_addr *net.Conn, message string) {
	for _connection, _ := range(*connections_addr) {
		if _connection != *connection_addr {
			fmt.Fprintf(_connection, message + "\n")
		}
	}
}

func get_message(conn *net.Conn) (string, error) {
	message, err := bufio.NewReader(*conn).ReadString('\n')
  	return strings.TrimRight(string(message), "\r\n"), err
}
