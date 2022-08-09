package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc := conn()
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	done := make(chan bool)
	go forever()
	<-done // Block forever

}

func conn() *nats.Conn {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		fmt.Println(err)
	}
	return nc
}

func forever() {
	for {
		// fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
