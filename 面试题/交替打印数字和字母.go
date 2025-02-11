package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers, letters := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-numbers:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				letters <- true
			}
		}
	}()
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
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				numbers <- true
			}
		}
	}(&wait)
	numbers <- true
	wait.Wait()
}
