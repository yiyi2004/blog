package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch := make(chan *nats.Msg, 10)

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println(string(msg.Data))
			default:
			}
		}
	}()

	_, err = conn.ChanSubscribe("demon", ch)
	if err != nil {
		log.Fatal(err)
	}
	conn.Flush()

	select {}
}
