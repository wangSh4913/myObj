package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
通过waitGroup包中方法科学等待协程退出
*/

//声明一个全局的锁
var wg sync.WaitGroup

func f1(i int) {
	//完成后会将全局协程数量减1
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(11) // [0,11)
	fmt.Printf("goroutine %d begin sleep %ds...\n",i,n)
	time.Sleep(time.Second * (time.Duration(n))) //随机睡几秒
}

func main() {
	fmt.Println(time.Now())
	for i := 0; i <= 10; i++ {
		//调用前把全局的协程数量加1
		wg.Add(1)
		//启动一个协程
		go f1(i)
		time.Sleep(1)
	}
	//这里会等待全局的所有协程退出后主协程才退出
	wg.Wait()
	fmt.Println("all goroutine exit")
	fmt.Println(time.Now())
}
