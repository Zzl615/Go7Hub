package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(name, "监控退出，停止了...", ctx.Value(key))
			return
		default:
			//取出值
			fmt.Println(name, "goroutine监控中...", ctx.Value(key))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	var value string = "上下文传值"
	valueCtx := context.WithValue(ctx, key, value)
	go watch(valueCtx, "【监控1】")
	go watch(valueCtx, "【监控2】")
	go watch(valueCtx, "【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
