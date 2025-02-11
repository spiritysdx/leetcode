package main

import (
    "fmt"
    "sync"
)

func main() {
    numbers, letters := make(chan bool), make(chan bool)
    wait := sync.WaitGroup{}
    
    // 打印数字的 goroutine
    go func() {
        i := 1
        for {
            select {
            case <-numbers:
                fmt.Println(i)
                i++
                letters <- true
            }
        }
    }()
    
    // 打印字母的 goroutine
    wait.Add(1)
    go func(wait *sync.WaitGroup) {
        i := 'A'
        for {
            select {
            case <-letters:
                if i > 'Z' {
                    wait.Done()
                    return
                }
                fmt.Println(string(i))
                i++
                numbers <- true
            }
        }
    }(&wait)
    
    // 启动第一个数字的打印
    numbers <- true
    wait.Wait()
}
