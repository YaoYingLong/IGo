package parallel

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
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
		//time.Sleep(2 * time.Second)
		ch <- "roger"
	}
}

func Server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen:", address)

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go HandleSession(conn, exitChan)
	}
}

func HandleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")
	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)

			if !ProcessTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}

			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}

func ProcessTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")

		exitChan <- 0
		return false
	}

	fmt.Println(str)

	return true
}
