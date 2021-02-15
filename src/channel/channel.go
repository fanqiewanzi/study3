package channel

import (
	"fmt"
	"sync"
	"time"
)

//定义一个锁的变量
var mutex sync.Mutex

//第一个goroutine
func print1(ch chan int) {
	fmt.Print("goroutine1\t")
	fmt.Println(<-ch)
}

//第二个goroutine
func print2(ch chan int) {

	fmt.Print("goroutine2\t")
	fmt.Println(<-ch)

}
func PrintTest() {
	//创建一个能容纳100个int类型的管道
	chs := make(chan int, 1)
	for i := 1; i <= 100; i++ {
		chs <- i
		//交给goroutine进行打印输出
		go print1(chs)
		i++
		//给管道赋值
		chs <- i
		//交给goroutine进行打印输出
		go print2(chs)
	}
	//等待两秒等所有goroutine运行完
	time.Sleep(2 * time.Second)
}
