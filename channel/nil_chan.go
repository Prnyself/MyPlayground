package channel

import (
	"errors"
	"fmt"
	"sync"
)

func nilChan(ch chan<- struct{}) func() {
	return func() {
		if ch != nil {
			fmt.Println("channel not nil")
			ch <- struct{}{}
		}
		fmt.Println("before sending to channel")
		fmt.Println("after sending to channel")
	}
}

func ShowNilChan() error {
	ch := make(chan struct{})
	// var ch chan struct{}
	go nilChan(ch)()
	<-ch
	fmt.Println("about to return")
	return nil
}

func nilWg(wg *sync.WaitGroup) func() {
	return func() {
		fmt.Println("before done")
		if wg != nil {
			wg.Done()
		}
		fmt.Println("after done")
	}
}

func ShowNilWg() error {
	var wg sync.WaitGroup
	wg.Add(1)
	go nilWg(&wg)()
	wg.Wait()
	fmt.Println("about to return")
	a := errors.New("abc")
	b := fmt.Errorf("panice: xxx, %w", a)
	// pkg/errors
	errors.Is(b, a)
	return nil
}
