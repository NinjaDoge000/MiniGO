package main

import (
	"fmt"
	"sync"
	"time"
)

const numGoods int = 10

var goodNum int

func producer(goods chan<- struct{}, wg *sync.WaitGroup) {

	defer wg.Done()
	defer close(goods)

	// produces 100 goods
	for i := 0; i < numGoods; i++ {
		goods <- struct{}{}
		fmt.Println("Produced good ", i+1)
	}

	// simulate work
	time.Sleep(100 * time.Millisecond)
}

func consumer(num int, goods <-chan struct{}, wg *sync.WaitGroup, mutex *sync.Mutex) {

	defer wg.Done()

	for range goods {

		mutex.Lock()
		goodNum++
		mutex.Unlock()

		fmt.Println("consumer ", num, "consuming good ", goodNum)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	// common channel
	goods := make(chan struct{}, numGoods)

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	wg.Add(4)
	go producer(goods, &wg)
	go consumer(1, goods, &wg, &mutex)
	go consumer(2, goods, &wg, &mutex)
	go consumer(3, goods, &wg, &mutex)

	wg.Wait()

}
