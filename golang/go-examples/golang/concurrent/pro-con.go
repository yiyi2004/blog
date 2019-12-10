package concurrent

import "fmt"

// Producer -
func Producer(facter int, out chan int) {
	for i := 0; i < 200000; i++ {
		out <- i * facter
	}
}

// Consumer -
func Consumer(in chan int) {
	for val := range in {
		fmt.Println(val)
	}
}
