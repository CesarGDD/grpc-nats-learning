package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc := conn()
	nc.Publish("foo", []byte("Hello World"))
	fmt.Println("hello publish")
}

func conn() *nats.Conn {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		fmt.Println(err)
	}
	return nc
}
