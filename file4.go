package main

import "fmt"

func Code4() {

	c := make(chan int)

	go func() {
		c <- 42

	}()
	fmt.Println(<-c)

	// directional

	cr := make(<-chan int)
	cs := make(chan<- int)
	fmt.Printf("%T, %T", cr, cs)

}
