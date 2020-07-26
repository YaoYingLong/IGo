package test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestStrings(t *testing.T) {
	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Atoi("10"))

	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(20, 2))

}

func coverPanic() {
	message := recover()
	fmt.Println("panic msg:", message)
}

func TestPanicRecover(t *testing.T) {
	defer coverPanic()
	panic("I am panic")
	panic(errors.New("kkkk"))
}
