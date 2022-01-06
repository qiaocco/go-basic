package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()

	time.Sleep(time.Second)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
