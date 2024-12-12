package main

import "fmt"

func main() {
	channel := make(chan int, 100)
	go setList(channel)

	for v := range channel {
		fmt.Println("recebendo: ", v)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("enviando: ", i)
	}
	close(channel)
}
