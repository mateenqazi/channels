package main

import "fmt"

func Code4() {

	c := make(chan int)
	fmt.Println(<-c)
	// c <- 42
	// go func() {
	// 	// c <- 42
	// 	fmt.Println(<-c)

	// }()
}
