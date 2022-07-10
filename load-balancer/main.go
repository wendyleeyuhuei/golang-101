package main

import (
	"container/heap"
	"math/rand"
	"time"
)

const (
	nRequester = 100
	nWorker    = 10
)

type Request struct {
	fn func() int // The operation to perform.
	c  chan int   // The channel to return the result.
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

type Pool []*Worker

/* Make Pool an implementation of heap interface */

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *Pool) Push(x any) {
	// dereference the pointer to assign value to it
	*p = append(*p, x.(*Worker)) // type assertion on x
}

func (p *Pool) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

type Balancer struct {
	pool Pool         // a pool of workers
	done chan *Worker // for requester to report task completion
}

// requester continuously sends request to work channel and further process with the request's result
func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		work <- Request{operation, c} // send request
		result := <-c                 // wait for answer
		furtherProcess(result)
	}
}

// work continuously deals with the request and inform done channel when it's finished
func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}

// balance dispatches the request to worker and updates a worker's load tracking data
func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

// dispatch sends request to the least loaded worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// Increment the queue
	w.pending++
	// Push back to heap.
	heap.Push(&b.pool, w)
}

// completed updates heap when job is finished
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// operation returns a random integer
func operation() int {
	rand.Seed(time.Now().UnixNano()) // yields a constantly-changing number
	n := rand.Int63n(int64(time.Second))
	time.Sleep(time.Duration(nWorker * n))
	return int(n)
}

func furtherProcess(result int) {
	// Kill some time (fake load).
	time.Sleep(time.Duration(rand.Intn(result)) * time.Millisecond)
}

func newBalancer() *Balancer {
	// Create a done channel
	done := make(chan *Worker, nWorker)
	// Create a balancer
	b := &Balancer{make(Pool, 0, nWorker), done}
	for i := 0; i < nWorker; i++ {
		// Create worker with requests buffered channel
		w := &Worker{requests: make(chan Request, nRequester), index: i}
		// Push the worker to heap
		heap.Push(&b.pool, w)
		// Launch work in goroutine
		go w.work(b.done)
	}
	return b
}

func main() {
	work := make(chan Request)
	for i := 0; i < nRequester; i++ {
		go requester(work)
	}
	newBalancer().balance(work)
}
