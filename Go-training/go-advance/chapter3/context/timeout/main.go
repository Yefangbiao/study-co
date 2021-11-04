package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	handle(context.Background(), 1*time.Second)
}

// handle handle模拟了一个超时控制的处理。
// 有两种情况返回 1.超时 2.得到数据
func handle(ctx context.Context, duration time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	// 注意这里防止内存泄漏缓冲区为1
	r := make(chan int, 1)
	go func() {
		timeConsumingFunc(r)
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("timeout: %v ms, context exit: %+v\n", duration, ctx.Err())
	case res := <-r:
		fmt.Printf("result: %d", res)
	}
}

func timeConsumingFunc(r chan<- int) {
	time.Sleep(3 * time.Second)
	r <- 1
}
