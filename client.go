package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide IPAddress:Port, you only provided one... try 'go run client.go 127.0.0.1:1234'")
		return
	}

	PORT := args[1]
	c, err := net.Dial("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Connected to server on: %v\n", args[1])
	fmt.Print("Type EXIT to end program\n")

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter your quote selection (1-15 or r for a random quote) or type \"users\" to see the number of connected clients")
		fmt.Print("$ ")
		text, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(c, text+"\n")

		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(message)
		if strings.TrimSpace(string(text)) == "EXIT" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
