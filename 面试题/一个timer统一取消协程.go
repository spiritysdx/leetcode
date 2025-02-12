package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func test(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Println("Running")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	timer := time.AfterFunc(5*time.Second, func() {
		fmt.Println("全局取消")
		cancel()
	})
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go test(ctx, &wg)
	}
	wg.Wait()
	timer.Stop()
	fmt.Println("All Done")
}
