# Golang 101

## [Recap] Goroutines

- A goroutine is an independently executing function, launched by a `go` statement.
- It has its own call stack, which grows and shrinks as required.
- It's not a thread.

### Demo

[demo 1](https://github.com/wendyleeyuhuei/golang-101/tree/main/goroutine-1)

[demo 2](https://github.com/wendyleeyuhuei/golang-101/tree/main/goroutine-2)

[demo 3](https://github.com/wendyleeyuhuei/golang-101/tree/main/goroutine-3)

- What's missing?
  - The main function couldn't see the output from the other goroutine.
  - No communication at all.

## Go Channel

### Definition

- Provides a **connection** between two goroutines, allowing them to communicate. 
- **Bidirectional** as default. Goroutines can send/receive data through the same channel. 
  - Can be single-directional i.e. send-only or receive-only channel 
- **First-class values**, just like strings or integers. 
- A channel can only transfer values of the **same type**.

