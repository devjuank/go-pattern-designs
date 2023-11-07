package main

import (
	"fmt"
	"sync"
)

var balance = 100

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()

	lock.Lock()
	balance = balance + amount
	lock.Unlock()
}

func Balance() int {
	return balance
}
func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())
}
