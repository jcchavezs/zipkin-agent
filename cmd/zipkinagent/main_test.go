package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
)

func TestTCPConnection(t *testing.T) {
	message := `
	[
	  {
		"traceId": "5af7183fb1d4cf5f",
		"id": "352bff9a74ca9ad2",
		"kind": "CLIENT",
		"debug": false,
		"shared": false,
		"tags": {
		  "sql.query": "SELECT * FROM users"
		}
	  }
	]
`

	go func() {
		conn, err := net.Dial("tcp", ":9412")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
		}
	}()

	l, err := net.Listen("tcp", ":9412")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		}

		conn.Close()

		return
	}

}
