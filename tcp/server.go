package main

import (
	"bufio"
	"fmt"
	"github.com/tomohiko-ito-10antz/go-dev/fizzbuzz"
	"log"
	"net"
)

func main() {
	tcpAddr := &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 1234,
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf(`Fail to listen: %+v`, err)
	}
	log.Println("Listen")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf(`Fail to accept: %+v`, err)
		}
		log.Println("Accept")
		go func(conn net.Conn) {
			defer conn.Close()
			scanner := bufio.NewScanner(conn)
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				log.Println("Receive")
				_, err := fmt.Fprintln(conn, fizzbuzz.FizzBuzz(scanner.Text()))
				if err != nil {
					log.Printf(`Fail to send: %+v`, err)
					return
				}
				log.Println("Send")
			}
		}(conn)
	}
}
