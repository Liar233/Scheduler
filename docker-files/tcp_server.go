package main

import (
	"os"
	"log"
	"strconv"
	"net"
	"fmt"
	"bufio"
)

func main() {
	port := getPort()

	startTCPServer(port)
}

func startTCPServer(port string) {
	server, err := net.Listen("tcp4", port)

	if err != nil {
		log.Fatalln(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Panicln(err)
			return
		}

		go handler(conn)
	}
}

func handler(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}


		log.Printf("%s", netData)
		c.Close()
}

func getPort() string {
	if len(os.Args) == 1 {
		log.Fatalln("Port not set...")
	}

	if _, err := strconv.ParseUint(os.Args[1], 10, 64); err != nil {
		log.Fatalln("Port in not numeric...")
	}

	return ":" + os.Args[1]
}
