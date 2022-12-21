package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		for range time.Tick(2 * time.Second) {
			n := rand.Intn(10)
			fmt.Println("generated number:", n)
			ch1 <- n
		}
	}()

	go func() {
		n := []int{}

		for {
			n = append(n, <-ch1)

			if len(n) == 3 {
				for _, nn := range n {
					fmt.Printf("number: %d\n", nn)
				}
				fmt.Println("")
				n = []int{}
			}
		}
	}()

	select {}
}
