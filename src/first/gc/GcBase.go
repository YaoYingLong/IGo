package gc

import (
	"log"
	"runtime"
	"time"
)

/*
Go语言自带垃圾回收机制（GC）,GC 是自动进行的,可以使用 runtime.GC() 函数手动 GC

finalizer（终止器）是与对象关联的一个函数，通过 runtime.SetFinalizer 来设置，
如果某个对象定义了 finalizer，当它被 GC 时候，这个 finalizer 就会被调用
*/

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}

func entry() {
	var rd = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
}

func GcBase() {
	entry()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
