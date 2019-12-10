package concurrent

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	ctx := context.Background()

	selected, selectErrors := generator(100).Select(ctx, 10, func(wid int, str string) bool {
		values := strings.Split(str, "\t")
		time.Sleep(50 * time.Millisecond)
		return len(values)%2 == 0
	})

	log.Println("This is printed before Select finishes because Select runs in the background")

	mapped, mapErrors := selected.Map(ctx, 10, func(wid int, str string) string {
		values := strings.Split(str, "\t")
		time.Sleep(50 * time.Millisecond)
		return strings.Join(values, ",")
	})

	for item := range mapped {
		log.Println(item)
	}

	// This is how we wait for Select to finish
	for err := range selectErrors {
		log.Println("There was an error:", err)
	}

	// This is how we wait for Map to finish
	for err := range mapErrors {
		log.Println("There was an error:", err)
	}
}
