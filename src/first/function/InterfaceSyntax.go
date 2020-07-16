package function

import (
	"fmt"
	"reflect"
)

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

type FuncCallerV2 func(interface{})

func (f FuncCallerV2) Call(p interface{}) {
	fmt.Println("form function:", p)
}

type DataWriter interface {
	WriteData(data interface{}) error

	CanWriter() bool
}

type Closer interface {
	WriterClose() error
}

type DataWriteCloser interface {
	DataWriter
	Closer
}

type FileWriter struct {
}

type NetWriter struct {
	FileWriter
}

func (d *FileWriter) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

func (d *FileWriter) WriterClose() error {
	fmt.Println("FileWriter Close")
	return nil
}

func (d *NetWriter) CanWriter() bool {
	fmt.Println("CanWrite")
	return false
}

type insImpl struct {
}

func (ins *insImpl) String() string {
	return "hi"
}

func GetStringer() fmt.Stringer {
	var ins *insImpl = nil
	return ins
}

func InterBaseFunc() {
	var invoker Invoker = new(Struct)
	invoker.Call("hello1")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("form function:", v)
	})
	invoker.Call("hello2")

	invoker = new(FuncCallerV2)
	invoker.Call("hello3")

	fileWriter := new(FileWriter)
	fileWriter.WriterClose()
	fileWriter.WriteData("FileWriter")

	netWriter := new(NetWriter)
	netWriter.WriterClose()
	netWriter.WriteData("NetWriter")
	netWriter.CanWriter()

	f := new(NetWriter)
	var writer DataWriter
	writer = f
	writer.WriteData("data")
	writer.CanWriter()

	var writer2 DataWriter = new(NetWriter)
	writer2 = f
	writer2.WriteData("data2")
	writer2.CanWriter()

	fmt.Println("GetStringer() == nil:", GetStringer() == nil)
	fmt.Println("GetStringer() type", reflect.TypeOf(GetStringer()))

	var ins *insImpl = nil
	fmt.Println("ins == nil:", ins == nil)
	fmt.Println("ins type", reflect.TypeOf(ins))

	var writeCloser DataWriteCloser = new(NetWriter)
	writeCloser.WriterClose()
	writeCloser.WriteData("writeCloser")
	writeCloser.CanWriter()
}
