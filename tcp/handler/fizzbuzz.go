package handler

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func FizzBuzz(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)
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
		_, err = fmt.Fprintln(conn, result)
		if err != nil {
			log.Printf("Error %+v\n", err)
			return
		}
		log.Printf("Send %s\n", result)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf(`Fail to read: %+v`, err)
	}
	log.Printf("Close")
}
