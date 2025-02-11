package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个生产者消费者模型，生产者有三个，分别是A、B、C，消费者只有一个D，A、B的生产速率为1、C的生产速率为2，且B、C在生产完5个消息后退出、A在生产完10个消息后退出。
// D在消费消息的时候是完全平等的。

func main() {
	wg := sync.WaitGroup{}
	var a, b, c bool
	numbers := make(chan int, 8)
	// A
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer func() {
			a = true
			wg.Done()
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			numbers <- i
		}
	}(&wg)
	// B
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer func() {
			b = true
			wg.Done()
		}()
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			numbers <- i
		}
	}(&wg)
	// C
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer func() {
			c = true
			wg.Done()
		}()
		for i := 1; i <= 5; i++ {
			time.Sleep(5000 * time.Millisecond)
			numbers <- i
		}
	}(&wg)
	// D
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer func() {
			wg.Done()
		}()
		for {
			n, _ := <-numbers
			if n != 0 {
				fmt.Println(n)
			}
			if a && b && c {
				break
			}
		}
	}(&wg)
	wg.Wait()
}
