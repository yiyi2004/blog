package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
)

var (
	TOPICS = os.Getenv("TOPICS")
)

func main() {
	var (
		wg  = new(sync.WaitGroup)
		seq = 0
	)

	topics := strings.FieldsFunc(TOPICS, func(r rune) bool {
		return r == ','
	})

	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, topic := range topics {
		for i := 0; i < 20; i++ {
			wg.Add(1)
			seq++
			go func() {
				conn.Publish(topic, []byte(fmt.Sprintf("[%d:%s]: hello", seq, topic)))
				conn.Flush()
				wg.Done()
			}()
		}
	}

	wg.Wait()
}
