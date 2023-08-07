package main

import (
	"fmt"
	"math/rand"
)

func Code() {
	channel := make(chan string)
	numsRounds := 3

	go funcThrow(channel, numsRounds)

	for i := 0; i < numsRounds; i++ {

		fmt.Println("here ", i, " : ", <-channel)
	}
}

func funcThrow(channel chan string, numsRounds int) {
	for i := 0; i < numsRounds; i++ {
		score := rand.Intn(10)
		fmt.Println("function inside the funcThrow")
		channel <- fmt.Sprint("your score ", score)
	}
}
