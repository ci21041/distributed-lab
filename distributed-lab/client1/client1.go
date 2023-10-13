package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	for {
		fmt.Println("Enter your message: ")
		stdin := bufio.NewReader(os.Stdin)
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintf(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting to server")
		return
	}
	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	for {
		write(conn)
	}
	//TODO Start getting and sending user messages.
}
