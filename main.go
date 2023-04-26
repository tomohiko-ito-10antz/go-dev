package main

import (
	"log"
	"os"
	"strconv"

	"github.com/tomohiko-ito-10antz/go-dev/tcp"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "server":
		tcp.FizzBuzzServer()
	case "client":
		target := os.Args[2]
		n, err := strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			log.Fatalf("Fail to parse args %v: %+v", os.Args, err)
		}
		tcp.FizzBuzzClient(target, n)
	case "client2":
		target := os.Args[2]
		n, err := strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			log.Fatalf("Fail to parse args %v: %+v", os.Args, err)
		}
		tcp.FizzBuzzClient2(target, n)
	}
}
