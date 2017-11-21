package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/jcchavezs/zipkin-agent"
	"github.com/jcchavezs/zipkin-agent/transport"
)

const (
	host           = "localhost"
	port           = "9412"
	connectionType = "tcp"
)

func main() {
	l, err := net.Listen(connectionType, host+":"+port)
	if err != nil {
		fmt.Printf("Error listening: %s.\n", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Printf("Listening on %s:%s\n", host, port)

	tn := os.Getenv("TRANSPORT")

	t, err := getTransporter(tn)
	if err != nil {
		fmt.Errorf("Failed when initializing the transport: %s\n", err.Error())
		os.Exit(1)
	}

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
func getTransporter(transportName string) (zipkinagent.Transporter, error) {
	switch transportName {
	case "logger":
		fmt.Printf("Sending to logs\n")
		return zipkinagent.NewLoggerTransporter(), nil
	default:
		url := os.Getenv("TRANSPORT_HTTP_URL")
		if url == "" {
			url = "http://localhost:9411/api/v2/spans"
		}

		fmt.Printf("Sending over http to endpoint %s\n", url)
		return transport.NewHttpTransporter(url), nil
	}
}

func handleRequest(conn net.Conn, c *zipkinagent.Collector) {
	defer conn.Close()

	spans := []zipkinagent.Span{}

	d := json.NewDecoder(conn)
	if err := d.Decode(&spans); err != nil {
		fmt.Printf("Failed to decode incoming data: %s\n", err.Error())
	}

	c.Collect(&spans)
}
