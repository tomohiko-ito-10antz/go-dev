package tcp

import (
	"log"
	"net"

	"github.com/tomohiko-ito-10antz/go-dev/tcp/handler"
)

func FizzBuzzServer() {
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
			handler.FizzBuzz(conn)
			conn.Close()
		}(conn)
	}
}
