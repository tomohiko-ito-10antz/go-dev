package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf(`Fail to parse args: %v`, os.Args)
	}
	address := os.Args[1]
	n, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf(`Fail to dial: %+v`, err)
	}

	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Fatalf(`Fail to dial: %+v`, err)
	}

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)
	for i := int64(0); i < n; i++ {
		_, err := fmt.Fprintln(conn, strconv.FormatInt(i, 10))
		if err != nil {
			log.Fatalf(`Fail to send %d: %+v`, i, err)
		}
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf(`Fail to read: %+v`, err)
		}
	}
}
