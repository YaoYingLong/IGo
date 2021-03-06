package main

import (
	"first/other"
	"first/syntax"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 基本语法
	syntax.BaseFunc()
	// 并行编程
	//parallel.ModelExec()
	// web服务
	//web.CusService()
	// 流程控制
	//control.BaseFunc()
	// 容器
	//container.BaseFunc()
	// 函数定义
	//function.BaseFunc()
	// 接口
	//function.InterBaseFunc()
	// 结构体
	//function.StructBase()
	rand.Seed(time.Now().UTC().UnixNano())
	other.Lissajous(os.Stdout)
}
