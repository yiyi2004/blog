package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		go func() {
			for {
				go func(ch chan int) {
					fmt.Printf("1:%v\n", <-ch)
				}(ch)

			}
		}()

		go func() {
			for {
				go func(ch chan int) {
					fmt.Printf("2:%v\n", <-ch)
				}(ch)
			}
		}()

		go func() {
			for {
				go func(ch chan int) {
					fmt.Printf("3:%v\n", <-ch)
				}(ch)
			}
		}()

	}()

	for i := 0; i < 40; i++ {
		ch <- i
	}

	time.Sleep(time.Second)
}
