package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("calulate ExpensiveFibonacci for %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InPogress map[int]bool
	IsPending map[int][]chan int
	Lock      sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exists := s.InPogress[job]
	if exists {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received: %d\n", resp)
	}

	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InPogress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate Fibonacci for %d\n", job)
	result := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready job: %d\n", job)
	}

	s.Lock.Lock()
	s.InPogress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InPogress: make(map[int]bool),
		IsPending: make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}

	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
