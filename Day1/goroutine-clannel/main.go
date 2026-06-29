package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		//存入数据
		ch <- i
		fmt.Printf("生成者发送：%d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		//读取数据
		fmt.Printf("消费者收到：%d\n", v)

	}
}

func main() {
	ch := make(chan int, 2) //带缓冲，容量 2
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "来自 ch1"
	ch2 <- "来自 ch2"

	for i := 0; i < 2; i++ {
		//阻塞通道直到通道接受到数据
		select {
		case msg := <-ch1:
			fmt.Println("select 收到：", msg)
		case msg := <-ch2:
			fmt.Println("select 收到：", msg)
		}
	}
}
