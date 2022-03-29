package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		println("go 1 start")
		ch <- 1
		println("go 1 end")
	}()
	go func() {
		defer wg.Done()
		println("go 2 start")
		ch <- 2
		println("go 2 end")
	}()
	go func() {
		wg.Wait()
		close(ch)
		println("close chan")
	}()
	for item := range ch {
		println("ans : " + fmt.Sprintf("%d", item))
	}
	panic(fmt.Errorf("this is err"))
}
