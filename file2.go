package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i // Send value to channel
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Consuming %d\n", value)
		time.Sleep(1 * time.Second) // Simulate some processing time
	}
}

func Code2() {
	bufferSize := 3
	ch := make(chan int, bufferSize) // Create a buffered channel

	go producer(ch)
	consumer(ch)
}
