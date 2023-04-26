package main

import (
	"log"
	"os"
	"strconv"

	"github.com/tomohiko-ito-10antz/go-dev/tcp"
)

func main() {
	address := os.Args[1]
	n, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf("Fail to parse args %v: %+v", os.Args, err)
	}

	tcp.FizzBuzzClient(address, n)
}
