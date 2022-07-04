package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("hi")
	//fmt.Println("I'm listening.")
	//time.Sleep(2 * time.Second)
	//fmt.Println("You're boring; I'm leaving.")
}

// Prints a message along with index infinitely,
// and each loop sleeps for random millisecond
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
