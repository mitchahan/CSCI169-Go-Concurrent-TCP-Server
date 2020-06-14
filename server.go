package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

// COUNT: to handle the number of clients connected to the concurrent TCP server
var count = 0

// PRINTER: to help make the long else if statements cleaner, no need to write tons of times
func printer(s *bufio.Scanner, num int, c net.Conn) {
	x := 1
	for x <= num*2 {
		s.Scan()
		x++
	}
	s1 := s.Text() + "\n"
	c.Write([]byte(string(s1)))
}

func handleConnection(c net.Conn) {
	fmt.Println("Connected to a client")
	fmt.Println("Hit Ctrl+C to close server\n")

	for {
		f, err := os.Open("quotes.txt")
		if err != nil {
			log.Fatal(err)
		}

		s := bufio.NewScanner(f)
		err = s.Err()
		if err != nil {
			log.Fatal(err)
		}

		recievedText, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(recievedText))
		fmt.Println("received command from a client:", temp)
		if temp == "EXIT" {
			count--
			break
		}
		if temp == "users" {
			counter := strconv.Itoa(count) + "\n"
			c.Write([]byte(string(counter)))
		} else if temp == "r" {
			r := rand.Intn(14) + 1
			printer(s, r, c)
		} else if temp == "1" {
			num := 1
			printer(s, num, c)
		} else if temp == "2" {
			num := 2
			printer(s, num, c)
		} else if temp == "3" {
			num := 3
			printer(s, num, c)
		} else if temp == "4" {
			num := 4
			printer(s, num, c)
		} else if temp == "5" {
			num := 5
			printer(s, num, c)
		} else if temp == "6" {
			num := 6
			printer(s, num, c)
		} else if temp == "7" {
			num := 7
			printer(s, num, c)
		} else if temp == "8" {
			num := 8
			printer(s, num, c)
		} else if temp == "9" {
			num := 9
			printer(s, num, c)
		} else if temp == "10" {
			num := 10
			printer(s, num, c)
		} else if temp == "11" {
			num := 11
			printer(s, num, c)
		} else if temp == "12" {
			num := 12
			printer(s, num, c)
		} else if temp == "13" {
			num := 13
			printer(s, num, c)
		} else if temp == "14" {
			num := 14
			printer(s, num, c)
		} else if temp == "15" {
			num := 15
			printer(s, num, c)
		} else {
			e := "That is not a valid command, please enter a command" + "\n"
			c.Write([]byte(string(e)))
		}

		f.Close()
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		count++
		go handleConnection(c)
	}
}
