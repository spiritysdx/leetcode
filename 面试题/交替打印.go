package main

import (
	"fmt"
	"sync"
)

// 必须两个add，两个done，记得哪个大于限制后关闭另一个的通道
func main() {
	wg := sync.WaitGroup{}
	ch1, ch2 := make(chan bool), make(chan bool)
	wg.Add(2) // 为两个goroutine都添加计数
	// 第一个goroutine
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		st := "Hello"
		i := 0
		for {
			select {
			case <-ch1:
				if i >= len(st) {
					close(ch2) // 关闭ch2
					return
				}
				fmt.Println(string(st[i]))
				i++
				ch2 <- true
			}
		}
	}(&wg)
	// 第二个goroutine
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		st := "Hello"
		i := 0
		for {
			select {
			case _, ok := <-ch2:
				if !ok { // 检查channel是否已关闭
					return
				}
				if i >= len(st) {
					close(ch1) // 关闭ch1
					return
				}
				fmt.Println(string(st[i]))
				i++
				ch1 <- true
			}
		}
	}(&wg)
	ch1 <- true // 启动第一个goroutine
	wg.Wait()
}
