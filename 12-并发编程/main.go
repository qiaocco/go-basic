package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func hello(i int) {
	fmt.Println("hello ", i)
	wg1.Done()
}
func main() { // 开启主goroutine去执行main函数

	wg1.Add(10000) // 计数牌
	for i := 0; i < 10000; i++ {
		go hello(i) // 打印顺序不固定
	}
	//go hello() // 开启goroutine去执行hello
	fmt.Println("hello main")
	//time.Sleep(time.Second)

	wg1.Wait() // 阻塞，等所有小弟干完活才收兵
}
