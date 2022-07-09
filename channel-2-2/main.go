package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			// do nothing
			case <-quit:
				// clean up
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}
