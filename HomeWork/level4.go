package main

import (
	"fmt"
)

func main() {
	over := make(chan bool, 1)
	//传入管道初始值，否则会锁住
	//我们我们把goroutine需要执行的内容放到外边来，就能让里边正常执行了，当完成遍历后，将over的值置为true即可
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 9 {
				over <- true
			}

		}
	}()
	<-over
	fmt.Println("over!!!")
}
