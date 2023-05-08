package main

import (
	"bytes"
	"fmt"
	"github.com/tomohiko-ito-10antz/go-dev/fizzbuzz"
	"log"
	"net"
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   nil,
		Port: 5678,
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf(`Fail to listen: %+v`, err)
	}
	log.Println("Listen")

	buffers := map[net.Addr]*bytes.Buffer{}
	var buf = make([]byte, 128)
	for {
		nRead, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatalf(`Fail to read: %+v`, err)
		}
		buffer, ok := buffers[addr]
		if !ok {
			buffer = bytes.NewBuffer(nil)
			buffers[addr] = buffer
		}
		_, err = buffer.Write(buf[:nRead])
		if err != nil {
			log.Fatalf(`Fail to read: %+v`, err)
		}
		for {
			var text string
			if _, err := fmt.Fscan(buffer, &text); err != nil {
				break
			}
			log.Println("Receive")
			_, err = conn.WriteTo([]byte(fizzbuzz.FizzBuzz(text)+"\n"), addr)
			if err != nil {
				log.Printf(`Fail to send: %+v`, err)
				break
			}
			log.Println("Send")
		}
	}
}
