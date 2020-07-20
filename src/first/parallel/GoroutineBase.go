package parallel

import (
	"errors"
	"fmt"
	"time"
)

func GoroutineBase() {
	ch1 := make(chan int)
	ch2 := make(chan interface{})
	type Equip struct {
		a int
		b int
	}
	ch3 := make(chan *Equip)

	ch1 <- 1
	ch2 <- "hi"
	ch3 <- &Equip{a: 1, b: 2}
	//fmt.Println("ch1:", ch1 -> )
}

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCService(ch chan string) {
	for {
		data := <-ch

		fmt.Println("server received:", data)
		ch <- "roger"
	}
}
