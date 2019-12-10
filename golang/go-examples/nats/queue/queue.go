package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	var (
		num1 int
		num2 int
	)

	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	conn.QueueSubscribe("demon", "worker1", func(m *nats.Msg) {
		num1++
		fmt.Printf("[first][%d]%s\n", num1, m.Data)
	})

	conn.QueueSubscribe("demon", "worker2", func(m *nats.Msg) {
		num2++
		fmt.Printf("[second][%d]%s\n", num2, m.Data)
	})

	conn.Subscribe("hello", func(m *nats.Msg) {
		fmt.Printf("[third][%d]%s\n", num2, m.Data)
	})

	runtime.Goexit()
}

// CheckErr -
func CheckErr(errChan chan error) {
	for err := range errChan {
		fmt.Printf("[err]: %s\n", err)
	}
}
