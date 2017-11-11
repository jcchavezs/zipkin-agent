package main

import (
	"encoding/json"
	"fmt"
	"github.com/jcchavezs/zipkin-agent"
	"net"
	"os"
)

const (
	host           = "localhost"
	port           = "3333"
	connectionType = "tcp"
)

func main() {
	l, err := net.Listen(connectionType, host+":"+port)
	if err != nil {
		fmt.Printf("Error listening: %s.\n", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + host + ":" + port)

	t := zipkinagent.NewLoggerTransporter()

	c, err := zipkinagent.NewCollector(t)
	if err != nil {
		fmt.Printf("Failed when initializing the collector: %s.\n", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %s.\n", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn, c)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, c *zipkinagent.Collector) {
	defer conn.Close()

	spans := []zipkinagent.Span{}

	d := json.NewDecoder(conn)

	if err := d.Decode(&spans); err != nil {
		fmt.Printf("Failed to decode incoming data: %s\n", err.Error())
	}

	c.Collect(&spans)
}
