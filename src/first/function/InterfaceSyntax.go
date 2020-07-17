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
	fmt.Println()

	invoker = new(FuncCallerV2)
	invoker.Call("hello3")
	fmt.Println()

	fileWriter := new(FileWriter)
	fileWriter.WriterClose()
	fileWriter.WriteData("FileWriter")
	fmt.Println()

	netWriter := new(NetWriter)
	netWriter.WriterClose()
	netWriter.WriteData("NetWriter")
	netWriter.CanWriter()
	fmt.Println()

	f := new(NetWriter)
	var writer DataWriter
	writer = f
	writer.WriteData("data")
	writer.CanWriter()
	fmt.Println()

	var writer2 DataWriter = new(NetWriter)
	writer2 = f
	writer2.WriteData("data2")
	writer2.CanWriter()
	fmt.Println()

	fmt.Println("GetStringer() == nil:", GetStringer() == nil)
	fmt.Println("GetStringer() type", reflect.TypeOf(GetStringer()))
	fmt.Println()

	var ins *insImpl = nil
	fmt.Println("ins == nil:", ins == nil)
	fmt.Println("ins type", reflect.TypeOf(ins))
	fmt.Println()

	var writeCloser DataWriteCloser = new(NetWriter)
	{
		rw, ok := writeCloser.(Closer)
		fmt.Println(ok, ";", rw.WriterClose())
	}
	{
		rw, ok := writeCloser.(*NetWriter)
		fmt.Println(ok, ";", rw.WriterClose(), rw.WriteData("*NetWriter"), rw.CanWriter())
	}
	writeCloser.WriterClose()
	writeCloser.WriteData("writeCloser")
	writeCloser.CanWriter()
	fmt.Println()
}
