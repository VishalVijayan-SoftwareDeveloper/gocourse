package Advanced

import (
	"fmt"
	"log"
	"net"
)

func TcpServer() {

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listner.Close()
	fmt.Println("Listening to port: 8080")

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println("accept error")
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	conn.Write([]byte("Hello from my TCp server !\n"))

}
