package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
使用context快速结束goroutine的执行
常用的方法：
   withCancel
   withDeadline 绝对的结束时间，完成工作的截止时间
   withTimeout  一段时间后结束，跟deadline差不多
   withValue  创建上下文，传值，往下传，特定场景使用
 */

var wg sync.WaitGroup

func f(ctx context.Context){
	defer wg.Done()
	LOOP:
	for  {
		select  {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(time.Second)
			fmt.Println("1111")
		}

	}
}

func main(){
	//创建一个context的对象，cancel是一个方法
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(10*time.Second)
	//调用cancel可以立即结束
	cancel()
	wg.Wait()
}
