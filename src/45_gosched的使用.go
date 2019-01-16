// 45_gosched的使用
package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
    runtime.GOMAXPROCS(1)  //使用单核
}

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go1")
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go3")
		}
	}()

	for i := 0; i < 5; i++ {
		//fmt.Println("world")
		fmt.Println("hello")
		if i == 3{
			runtime.Gosched()
		}
	}
	time.Sleep(time.Second*5)
}
