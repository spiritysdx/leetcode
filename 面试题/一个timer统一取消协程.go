package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // 监听全局 context 的取消信号
			fmt.Printf("Worker %d received stop signal, exiting...\n", id)
			return
		default:
			fmt.Printf("Worker %d is running...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) // 统一的 ctx 控制所有协程
	defer cancel()

	// 外部使用 Timer 触发取消
	timer := time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timeout reached! Stopping all workers...")
		cancel() // 触发全局取消
	})

	// 启动多个 worker
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	wg.Wait() // 等待所有 worker 退出
	timer.Stop() // 确保 timer 释放
	fmt.Println("All workers stopped, main function exits.")
}
