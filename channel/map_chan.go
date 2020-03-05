package channel

import (
	"fmt"
	"sync"
)

func MapChan() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan map[int]string)
	m := map[int]string{1: "a", 2: "b"}
	fmt.Println("before:", m)
	go func() {
		mc := <-ch
		mc[3] = "c"
		fmt.Println("in goroutine:", mc)
		wg.Done()
	}()
	ch <- m
	wg.Wait()
	fmt.Println("after:", m)
}
