package main

import "fmt"

func main() {
	P(4)
}

func testChannelClose() {
	ch := make(chan int, 1)

	ch <- 1
	close(ch)

	fmt.Println(<-ch)
}

func testMapLength() {
	m := make(map[int]string)

	m[1] = "demon"
	m[2] = "abser"
	m[3] = "caicai"

	fmt.Println(len(m))
}

// P -
func P(value int) {
	if value > 0 {
		P(value - 1)
		fmt.Printf("%d", value)
		P(value - 1)
	}
}

// think how we can use stack to implement this function
