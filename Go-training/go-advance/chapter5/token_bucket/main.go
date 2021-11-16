package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"math/rand"
	"sync"
	"time"
)

var errCount int

func main() {
	wg := &sync.WaitGroup{}
	bucket := newTokenBucket()
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			r := mockRequest()
			err := bucket.handler(r.ctx)
			if err != nil {
				fmt.Println("调用失败", i)
				errCount++
			} else {
				fmt.Println("调用成功", i)
			}
		}()
		time.Sleep(20 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("失败个数", errCount)
	fmt.Println(time.Since(start))
}

type request struct {
	ctx context.Context
}

func mockRequest() *request {
	// 模拟一个请求
	return &request{ctx: context.Background()}
}

type tokenBucket struct {
	limiter *rate.Limiter
}

func newTokenBucket() *tokenBucket {
	return &tokenBucket{
		limiter: rate.NewLimiter(rate.Every(time.Millisecond*25), 10),
	}
}

func (t *tokenBucket) handler(ctx context.Context) error {
	if t.limiter.Allow() {
		err := t.limiter.Wait(ctx)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("can not get token")
	}
	done := make(chan struct{}, 1)
	go func() {
		// 每次请求80-100ms
		wait := rand.Intn(20) + 80
		time.Sleep(time.Duration(wait) * time.Millisecond)
		defer close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return nil
}
