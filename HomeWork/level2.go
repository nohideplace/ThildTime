package main

import "sync"

var wg sync.WaitGroup
var x int

// 避免死锁
var ch = make(chan int, 1)

func Print1() {
	for x < 100 {
		x = <-ch
		println("打印偶数", x)
		x++
		ch <- x
	}
	wg.Done()
}
func Print2() {
	for x < 100 {
		x = <-ch
		println("打印奇数", x)
		x++
		ch <- x
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	ch <- 0
	//为啥这里放在前边的反而后执行啊
	go Print2()
	go Print1()
	wg.Wait()

}
