package main

import (
	"log"
	"os"
	"strconv"

	udp "github.com/tomohiko-ito-10antz/go-dev/udp"
)

func main() {
	address := os.Args[1]
	n, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf("Fail to parse args %v: %+v", os.Args, err)
	}

	udp.FizzBuzzClient(address, n)
}
