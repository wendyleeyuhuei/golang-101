package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create an unbuffered channel
	c := make(chan string)
	// Launch function in a goroutine
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}

	fmt.Println("You're boring; I'm leaving.")
}

// Prints a message along with index infinitely,
// and each loop sleeps for random millisecond
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
