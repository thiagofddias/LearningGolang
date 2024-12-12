package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	callDatabase(&wg)
	callAPI(&wg)
	processInternal(&wg)

	time.Sleep(3 * time.Second)
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finalizado callDatabase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("finalizado callAPI")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finalizado processInternal")
	wg.Done()
}
