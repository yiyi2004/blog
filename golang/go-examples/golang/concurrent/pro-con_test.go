package concurrent

import (
	"testing"
	"time"
)

func TestProducerConsumenr(t *testing.T) {
	ch := make(chan int)

	go Producer(2, ch)
	go Producer(5, ch)
	go Consumer(ch)

	time.Sleep(time.Second)
}
