package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func FizzBuzzClient(address string, n int64) {
	conn, err := net.Dial("tcp", address)
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

func FizzBuzzClient2(target string, n int64) {
	for i := int64(0); i < n; i++ {
		conn, err := net.Dial("tcp", target)
		if err != nil {
			log.Fatalf(`Fail to dial: %+v`, err)
		}
		scanner := bufio.NewScanner(conn)
		scanner.Split(bufio.ScanWords)

		_, err = fmt.Fprintln(conn, strconv.FormatInt(i, 10))
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
