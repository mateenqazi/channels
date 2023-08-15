package main

import "fmt"

func Code6() {

	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
	fmt.Println("about to exit")

}
