package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boring("hi")
}

// Prints a message along with index infinitely,
// and each loop sleeps for 1 second
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

// Prints a message along with index infinitely,
// and each loop sleeps for random millisecond
func slightlyLessBoring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
