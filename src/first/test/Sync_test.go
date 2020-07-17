package test

import (
	"first/parallel"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMutex(t *testing.T) {
	{
		for a, i := 0, 0; i < 1000; i++ {
			go func(idx int) {
				a += 1
				fmt.Println(a)
			}(i)
		}
	}
	time.Sleep(time.Second)
	fmt.Println()

	{
		var lock sync.Mutex
		for a, i := 0, 0; i < 1000; i++ {
			go func(idx int) {
				lock.Lock()
				defer lock.Unlock()
				a += 1
				fmt.Printf("goroutine %d, a=%d\n", idx, a)
			}(i)
		}
	}
	time.Sleep(time.Second)
}

func TestSyncRWMutex(t *testing.T) {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go parallel.Read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go parallel.Write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
