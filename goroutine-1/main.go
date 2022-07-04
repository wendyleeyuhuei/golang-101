package main

import (
	"fmt"
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
