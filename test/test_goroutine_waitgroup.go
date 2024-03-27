package main

import (
	"fmt"
	"sync"
)

func say2(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Printf("done\n")
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	say2("hello", &wg)
	say2("world", &wg)
	fmt.Println("over!")

}
