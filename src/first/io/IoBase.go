package io

import (
	"bufio"
	"bytes"
	"fmt"
)

func Read() {
	data := []byte("Go语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println("buf[:]:", string(buf[:]))
	fmt.Println("n, err:", n, err)
}

func ReadByte() {
	data := []byte("Go语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	n, err := r.ReadByte()
	fmt.Println("Result:", string(n), err)
}

func ReadBytes() {
	data := []byte("Go语言中文网, Go语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println("Result:", string(line), err)
}

func ReadLine() {
	data := []byte("Golang is a beautiful language. \r\n I like it!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println("result:", string(line), prefix, err)
}

func ReadRune() {
	data := []byte("Golang is a beautiful language. \r\n I like it!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)
}

func ReadSlice() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadSlice(delim)
	fmt.Println("result1:", string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println("result2:", string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println("result3:", string(line), err)
}

func ReadString() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadString(delim)
	fmt.Println("result1:", string(line), err)
	line, err = r.ReadString(delim)
	fmt.Println("result2:", string(line), err)
	line, err = r.ReadString(delim)
	fmt.Println("result3:", string(line), err)
}

func Buffered() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [14]byte
	n, err := r.Read(buf[:])
	fmt.Println("Read:", string(buf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println("Buffered:", rn)
	n, err = r.Read(buf[:])
	fmt.Println("Read2:", string(buf[:n]), n, err)
	rn = r.Buffered()
	fmt.Println("Buffered:", rn)
}

func Peek() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	bl, err := r.Peek(8)
	fmt.Println("Peek(8)", string(bl), err)
	bl, err = r.Peek(14)
	fmt.Println("Peek(14)", string(bl), err)
	bl, err = r.Peek(20)
	fmt.Println("Peek(20)", string(bl), err)
}

func Available() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("Go语言入门教程")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())
}

func WriteBuffered() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Buffered())
	w.Flush()
	fmt.Println("执行 Flush 方法后，写入的字节数为：", w.Buffered())
	fmt.Printf("执行 Flush 后缓冲区输出 %q\n", string(wr.Bytes()))
}

func Write() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	n, err := w.Write(p)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)
}

func WriteByte() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var c byte = 'G'
	err := w.WriteByte(c)
	w.Flush()
	fmt.Println(string(wr.Bytes()), err)
}

func WriteRune() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var r rune = 'G'
	size, err := w.WriteRune(r)
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)
}

func WriteString() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	s := "C语言中文网"
	n, err := w.WriteString(s)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)
}
