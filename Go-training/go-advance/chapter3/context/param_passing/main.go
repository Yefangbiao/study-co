package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// 传递参数
	ctx = WithParam(ctx, "hello", "world")
	// 取出参数
	value := ParamFromCtx(ctx, "hello")
	fmt.Println(value)
}

func WithParam(ctx context.Context, key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func ParamFromCtx(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}
