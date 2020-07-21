package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectType01(t *testing.T) {
	{
		var a int
		typeOfA := reflect.TypeOf(a)
		fmt.Println("name:", typeOfA.Name(), "kind:", typeOfA.Kind())
	}
	{
		type A struct {
		}
		a := &A{}
		typeOfA := reflect.TypeOf(a)
		fmt.Println("name:", typeOfA.Name(), "kind:", typeOfA.Kind())
		typeOfA = typeOfA.Elem()
		fmt.Println("name:", typeOfA.Name(), "kind:", typeOfA.Kind())
	}
	{
		type A struct {
		}
		a := A{}
		typeOfA := reflect.TypeOf(a)
		fmt.Println("name:", typeOfA.Name(), "kind:", typeOfA.Kind())
	}
	{
		type Enum int
		const Zero Enum = 0
		// 声明一个空结构体
		type cat struct {
		}
		// 获取结构体实例的反射类型对象
		typeOfCat := reflect.TypeOf(cat{})
		// 显示反射类型对象的名称和种类
		fmt.Println("name:", typeOfCat.Name(), "kind:", typeOfCat.Kind())
		// 获取Zero常量的反射类型对象
		typeOfA := reflect.TypeOf(Zero)
		// 显示反射类型对象的名称和种类
		fmt.Println("name:", typeOfA.Name(), "kind:", typeOfA.Kind())
	}
	{
		type cat struct {
			Name string
			Type int `json:"type" id:"100"`
		}
		ins := cat{Name: "mimi", Type: 1}
		typeOfCat := reflect.TypeOf(ins)
		for i := 0; i < typeOfCat.NumField(); i++ {
			fieldType := typeOfCat.Field(i)
			fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
		}

		if catType, ok := typeOfCat.FieldByName("Type"); ok {
			fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
		}
	}
	{
		type MyInt int
		var x MyInt = 7
		tof := reflect.TypeOf(x)
		fmt.Println("name:", tof.Name(), "kind:", tof.Kind())
		vof := reflect.ValueOf(x)
		ref := vof.Interface().(MyInt)
		fmt.Println("ref:", ref)
		fmt.Println("can set:", vof.CanSet())
		vof2 := reflect.ValueOf(&x)
		fmt.Println("can set:", vof2.CanSet())
		vof3 := vof2.Elem()
		fmt.Println("can set:", vof3.CanSet())
		vof3.SetInt(25)
		fmt.Println("ref:", vof3)
	}
}
