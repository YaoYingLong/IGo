package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func GetThingDone(param1, param2 int) {
	fmt.Println("param1, param2:", param1, param2)
}

func TestGoroutineCreate(t *testing.T) {
	go GetThingDone(1, 2)
	go func(param1, param2 int) {
		fmt.Println("param3, param4:", param1, param2)
	}(3, 4)
	go fmt.Println("param5, param6:", 5, 6)
	fmt.Println("param7, param8:", 7, 8)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick:", times)
		time.Sleep(time.Second)
	}
}

func TestGoroutineV2(t *testing.T) {
	go running()

	var input string
	fmt.Scanln(&input)
}

var count int

func incCount(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

func TestGoroutineV3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go incCount(&wg)
	go incCount(&wg)
	wg.Wait()
	fmt.Println("count:", count)
}

var counter int32

func incCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

func TestGoroutineV5(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go incCounter(1, &wg)
	go incCounter(2, &wg)
	wg.Wait()
	fmt.Println("counter:", counter)
}

var shutdown int64

func doWork(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
func TestGoroutineV6(t *testing.T) {
	println("cup count:", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	go doWork("A", &wg)
	go doWork("B", &wg)

	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

func TestChannel001(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan interface{})
	type Equip struct {
		a int
		b int
	}
	ch3 := make(chan *Equip)

	go func() { ch1 <- 1 }()
	go func() { ch2 <- "hi" }()
	go func() { ch3 <- &Equip{a: 1, b: 2} }()

	//data := <-ch1
	data, ok := <-ch1
	fmt.Println("data, ok:", data, ok)
	//fmt.Println("ch1:", <-ch1)
	fmt.Println("ch2:", <-ch2)
	fmt.Println("ch3:", <-ch3)

	go func() {
		for i := 3; i >= 0; i-- {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()
	for data := range ch1 {
		fmt.Println("data:", data)
		if data == 0 {
			break
		}
	}
}

func TestChannel002(t *testing.T) {
	ch := make(chan int)
	go func(ch chan int) {
		for {
			data := <-ch
			if data == 0 {
				break
			}
			fmt.Println("data:", data)
		}
		ch <- 0
	}(ch)

	for i := 1; i < 11; i++ {
		ch <- i
	}

	ch <- 0
	<-ch
}

func TestChannel003(t *testing.T) {
	ch := make(chan int)
	// 声明一个只能发送的通道类型
	var chSendOnly = make(chan<- int)
	chSendOnly <- 1
	//声明一个只能接收的通道类型
	var chRecvOnly = make(<-chan int)
	close(ch)
	close(chSendOnly)
	<-chRecvOnly
}

func Runner(baton chan int, wg *sync.WaitGroup) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton, wg)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

func TestChannel004(t *testing.T) {
	var wg sync.WaitGroup
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton, &wg)

	baton <- 1
	wg.Wait()
}

func player(name string, court chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}

func TestChannel005(t *testing.T) {
	court := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go player("Nadal", court, &wg)
	go player("Djokovic", court, &wg)

	court <- 1

	wg.Wait()
}

func TestChannel006(t *testing.T) {
	ch := make(chan int, 3)
	fmt.Println("before len:", len(ch))

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("after len:", len(ch))
	//ch<-4
	fmt.Println("ch:", <-ch)
	fmt.Println("after1 len:", len(ch))
	fmt.Println("ch:", <-ch)
	fmt.Println("after2 len:", len(ch))
	fmt.Println("ch:", <-ch)
	ch <- 4
	fmt.Println("after3 len:", len(ch))
	fmt.Println("ch:", <-ch)
}
