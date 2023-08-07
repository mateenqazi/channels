package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numWorkers = 4 // Number of worker goroutines

func squareWorker(id int, jobs <-chan int, results chan<- int) {
	for num := range jobs {
		result := num * num
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond) // Simulate some work
		results <- result
	}
}

func Code3() {

	// Create channels for sending and receiving jobs and results
	jobs := make(chan int)
	results := make(chan int)

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		go squareWorker(w, jobs, results)
	}

	// Generate a list of numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Send jobs to the workers
	go func() {
		for _, num := range numbers {
			jobs <- num
		}
		close(jobs)
	}()

	// Collect results from the workers
	var sumOfSquares int
	for i := 0; i < len(numbers); i++ {
		result := <-results
		sumOfSquares += result
	}

	fmt.Printf("List: %v\n", numbers)
	fmt.Printf("Sum of squares: %d\n", sumOfSquares)
}
