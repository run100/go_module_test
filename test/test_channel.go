package main

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {

	done := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("hello world")
		}
		done <- true
	}()

	<-done
	fmt.Println("over!")

	//println("111")
	//var ch chan int
	//ch := make(chan int)

	//chs := make([]chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go Count(chs[i])
	//}
	//
	//for _, ch := range chs {
	//	<-ch
	//}

	//chs := make([]chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go Count(chs[i])
	//}
	//
	//for _, ch := range chs {
	//	<-ch
	//}
}
