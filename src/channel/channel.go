package channel

import (
	"fmt"
	"time"
)

func print1(ch chan int) {
		fmt.Println(<-ch)
}
func print2(ch chan int)  {
	fmt.Println(<-ch)
}
func PrintTest()  {
	chs:=make(chan int,100)
	for i:=1;i<100;i++{
		chs<- i
		go print1(chs)
		go print2(chs)
	}
	time.Sleep(2 * time.Second) //延时2s
}