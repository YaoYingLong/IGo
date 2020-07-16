package test

import (
	"first/io"
	"testing"
)

func TestIoRead(t *testing.T) {
	io.Read()
}

func TestReadByte(t *testing.T) {
	io.ReadByte()
}

func TestReadBytes(t *testing.T) {
	io.ReadBytes()
}

func TestReadLine(t *testing.T) {
	io.ReadLine()
}

func TestReadRune(t *testing.T) {
	io.ReadRune()
}

func TestReadSlice(t *testing.T) {
	io.ReadSlice()
}

func TestReadString(t *testing.T) {
	io.ReadString()
}

func TestBuffered(t *testing.T) {
	io.Buffered()
}

func TestPeek(t *testing.T) {
	io.Peek()
}

func TestAvailable(t *testing.T) {
	io.Available()
}

func TestWriteBuffered(t *testing.T) {
	io.WriteBuffered()
}

func TestWrite(t *testing.T) {
	io.Write()
}

func TestWriteByte(t *testing.T) {
	io.WriteByte()
}

func TestWriteRune(t *testing.T) {
	io.WriteRune()
}

func TestWriteString(t *testing.T) {
	io.WriteString()
}
