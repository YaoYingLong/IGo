package parallel

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
