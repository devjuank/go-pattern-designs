package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Already connected")
	}

	return db
}

func (db *Database) CreateSingleConnection() {
	fmt.Println("Starting Connections")
	time.Sleep(2 * time.Second)
	fmt.Println("Connected")
}

func main() {
	var wg sync.WaitGroup
	count := 10
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
