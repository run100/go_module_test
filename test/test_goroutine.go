package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Add(x, y int) {
	z := x + y
	fmt.Printf("z is %v\n", z)
}

var counter int = 0

func count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Printf("counter++ is %v\n", counter)
	lock.Unlock()

}

func main() {

	lock := sync.Mutex{}

	for i := 0; i < 10; i++ {
		count(&lock)
	}

	//for i := 0; i < 10; i++ {
	//	go Add(i, 2)
	//}

	//time.Sleep(1000 * time.Millisecond)

	for {
		lock.Lock() // 上锁
		c := counter
		lock.Unlock() // 解锁

		runtime.Gosched() // 出让时间片

		if c >= 10 {
			break
		}
	}

}
