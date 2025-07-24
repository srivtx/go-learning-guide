package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker represents a worker in our worker pool
type Worker struct {
	ID   int
	Jobs <-chan Job
	Quit chan bool
}

// Job represents work to be done
type Job struct {
	ID   int
	Data string
}

// Result represents the result of a job
type Result struct {
	JobID  int
	Output string
	Worker int
}

// NewWorker creates a new worker
func NewWorker(id int, jobs <-chan Job) *Worker {
	return &Worker{
		ID:   id,
		Jobs: jobs,
		Quit: make(chan bool),
	}
}

// Start begins the worker's job processing
func (w *Worker) Start(results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case job := <-w.Jobs:
			// Simulate work
			time.Sleep(time.Millisecond * 100)
			result := Result{
				JobID:  job.ID,
				Output: fmt.Sprintf("Processed: %s", job.Data),
				Worker: w.ID,
			}
			results <- result
			fmt.Printf("Worker %d finished job %d\n", w.ID, job.ID)
			
		case <-w.Quit:
			fmt.Printf("Worker %d quitting\n", w.ID)
			return
		}
	}
}

// Counter demonstrates safe concurrent counter
type Counter struct {
	mu    sync.RWMutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely returns the counter value
func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// Pipeline stage functions
func generateNumbers(max int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return out
}

func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num%2 == 0 {
				out <- num
			}
		}
	}()
	return out
}

// Fan-out/Fan-in pattern
func fanOut(in <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		out := make(chan int)
		outputs[i] = out
		
		go func(output chan<- int) {
			defer close(output)
			for num := range in {
				// Simulate work
				time.Sleep(50 * time.Millisecond)
				output <- num * 2
			}
		}(out)
	}
	return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(input)
	}
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

// Timeout and cancellation example
func doWorkWithTimeout(timeout time.Duration) (string, error) {
	done := make(chan string, 1)
	
	go func() {
		// Simulate long-running work
		time.Sleep(2 * time.Second)
		done <- "Work completed!"
	}()
	
	select {
	case result := <-done:
		return result, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("work timed out after %v", timeout)
	}
}

func main() {
	fmt.Println("=== Advanced Go Concurrency Patterns ===")

	// ==================== BASIC GOROUTINES ====================
	fmt.Println("\n--- Basic Goroutines ---")
	
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 3; j++ {
				fmt.Printf("Goroutine %d: count %d\n", id, j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All basic goroutines finished")

	// ==================== CHANNELS ====================
	fmt.Println("\n--- Channel Communication ---")
	
	// Unbuffered channel
	messages := make(chan string)
	
	go func() {
		messages <- "Hello from goroutine!"
	}()
	
	msg := <-messages
	fmt.Printf("Received: %s\n", msg)
	
	// Buffered channel
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	
	fmt.Printf("Buffered channel contents: %d, %d, %d\n", <-buffered, <-buffered, <-buffered)

	// ==================== SELECT STATEMENT ====================
	fmt.Println("\n--- Select Statement ---")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received: %s\n", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}

	// ==================== WORKER POOL PATTERN ====================
	fmt.Println("\n--- Worker Pool Pattern ---")
	
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Start workers
	var workerWg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		workerWg.Add(1)
		worker := NewWorker(i, jobs)
		go worker.Start(results, &workerWg)
	}
	
	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{
			ID:   i,
			Data: fmt.Sprintf("job-data-%d", i),
		}
	}
	close(jobs)
	
	// Collect results
	go func() {
		workerWg.Wait()
		close(results)
	}()
	
	for result := range results {
		fmt.Printf("Result: Job %d -> %s (by Worker %d)\n", 
			result.JobID, result.Output, result.Worker)
	}

	// ==================== PIPELINE PATTERN ====================
	fmt.Println("\n--- Pipeline Pattern ---")
	
	// Set up pipeline: generate -> square -> filter even
	numbers := generateNumbers(10)
	squared := squareNumbers(numbers)
	evens := filterEven(squared)
	
	fmt.Print("Even squares: ")
	for result := range evens {
		fmt.Printf("%d ", result)
	}
	fmt.Println()

	// ==================== FAN-OUT/FAN-IN PATTERN ====================
	fmt.Println("\n--- Fan-out/Fan-in Pattern ---")
	
	input := generateNumbers(5)
	
	// Fan-out to multiple workers
	workers := fanOut(input, 3)
	
	// Fan-in results
	output := fanIn(workers...)
	
	fmt.Print("Fan-out/Fan-in results: ")
	for result := range output {
		fmt.Printf("%d ", result)
	}
	fmt.Println()

	// ==================== MUTEX FOR SHARED STATE ====================
	fmt.Println("\n--- Mutex for Shared State ---")
	
	counter := &Counter{}
	var counterWg sync.WaitGroup
	
	// Multiple goroutines incrementing counter
	for i := 0; i < 10; i++ {
		counterWg.Add(1)
		go func(id int) {
			defer counterWg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d finished incrementing\n", id)
		}(i)
	}
	
	counterWg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())

	// ==================== TIMEOUT AND CANCELLATION ====================
	fmt.Println("\n--- Timeout and Cancellation ---")
	
	// Test with short timeout (should timeout)
	result, err := doWorkWithTimeout(1 * time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n", result)
	}
	
	// Test with long timeout (should complete)
	result, err = doWorkWithTimeout(3 * time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n", result)
	}

	// ==================== DONE PATTERN ====================
	fmt.Println("\n--- Done Pattern ---")
	
	done := make(chan bool)
	
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Background goroutine stopping...")
				return
			default:
				fmt.Println("Background goroutine working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	done <- true
	time.Sleep(100 * time.Millisecond) // Give time for cleanup

	// ==================== CHANNEL DIRECTIONS ====================
	fmt.Println("\n--- Channel Directions ---")
	
	// Send-only channel function
	sender := func(ch chan<- string) {
		ch <- "Hello from sender"
		close(ch)
	}
	
	// Receive-only channel function
	receiver := func(ch <-chan string) {
		for msg := range ch {
			fmt.Printf("Received: %s\n", msg)
		}
	}
	
	ch := make(chan string)
	go sender(ch)
	receiver(ch)

	fmt.Println("\n=== Concurrency patterns demonstration complete ===")
}
