package tcp

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"strconv"
)

func FizzBuzzServer() {
	udpAddr := &net.UDPAddr{
		IP:   nil,
		Port: 5678,
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf(`Fail to listen: %+v`, err)
	}
	log.Println("Listen")
	buffers := map[*net.UDPAddr]*bytes.Buffer{}
	scanners := map[*net.UDPAddr]*bufio.Scanner{}
	var buf = make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalf(`Fail to read: %+v`, err)
		}
		if _, ok := buffers[addr]; !ok {
			buffers[addr] = bytes.NewBuffer(nil)
		}
		_, err = buffers[addr].Write(buf[:n])
		if err != nil {
			log.Fatalf(`Fail to read: %+v`, err)
		}

		if _, ok := scanners[addr]; !ok {
			scanners[addr] = bufio.NewScanner(buffers[addr])
			scanners[addr].Split(bufio.ScanWords)
		}

		scanner := scanners[addr]
		for scanner.Scan() {
			token := scanner.Text()
			log.Printf("Receive %s\n", token)

			var result string
			number, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				result = "Error"
			} else {
				switch {
				case number%15 == 0:
					result = "FizzBuzz"
				case number%3 == 0:
					result = "Fizz"
				case number%5 == 0:
					result = "Buzz"
				default:
					result = token
				}
			}
			_, err = conn.WriteToUDP([]byte(result+"\n"), addr)
			if err != nil {
				log.Printf("Error %+v\n", err)
				return
			}
			log.Printf("Send %s\n", result)
		}
	}
}
