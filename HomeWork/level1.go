package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int
var ch = make(chan int, 2)

//管道的类型实际上就是队列

// 我们这里将管道长度设置为2，避免过程阻塞
func add() {
	for i := 0; i < 50000; i++ {
		//获取隧道的值
		x = <-ch
		x++
		//传值入隧道
		ch <- x
	}
	wg.Done()

}

func main() {
	wg.Add(2)
	//首先要向管道传入值，否则函数里的取值也会死等
	ch <- 0
	go add()
	go add()
	wg.Wait()
	fmt.Println(<-ch)
}
