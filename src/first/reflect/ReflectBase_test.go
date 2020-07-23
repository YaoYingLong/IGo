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
	{
		fmt.Println()
		type T struct {
			A int
			B string
		}
		t := T{23, "skidoo"}
		s := reflect.ValueOf(&t).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
		fmt.Println("Field_0:", s.Field(0).CanSet())
		fmt.Println("Field_1:", s.Field(1).CanSet())
		s.Field(0).SetInt(99)
		s.Field(1).SetString("Sunset Strip")
		fmt.Println("t is now:", t)
	}
	{
		fmt.Println()
		type cat struct {
			Name string
			Type int `json:"type" id:"100"`
		}
		// 创建cat的实例
		ins := cat{Name: "mimi", Type: 1}
		// 获取结构体实例的反射类型对象
		typeOfCat := reflect.TypeOf(ins)
		// 遍历结构体所有成员
		for i := 0; i < typeOfCat.NumField(); i++ {
			// 获取每个成员的结构体字段类型
			fieldType := typeOfCat.Field(i)
			// 输出成员名和tag
			fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
		}
		// 通过字段名, 找到字段类型信息
		if catType, ok := typeOfCat.FieldByName("Type"); ok {
			// 从tag中取出需要的tag
			fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
		}
	}
}

func TestReflectValue01(t *testing.T) {
	{
		type dummy struct {
			a int
			b string
			float32
			bool
			next *dummy
		}
		// 值包装结构体
		d := reflect.ValueOf(dummy{
			next: &dummy{},
		})
		// 获取字段数量
		fmt.Println("NumField", d.NumField())
		// 获取索引为2的字段(float32字段)
		floatField := d.Field(2)
		// 输出字段类型
		fmt.Println("Field", floatField.Type())
		// 根据名字查找字段
		fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())
		// 根据索引查找值中, next字段的int字段的值
		fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
	}
	{
		fmt.Println()
		var a *int
		fmt.Println("var a *int is nil:", reflect.ValueOf(a).IsNil())
		fmt.Println("var a *int is valid:", reflect.ValueOf(a).IsValid())

		fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

		s := struct{}{}

		fmt.Println("not exist attr:", reflect.ValueOf(s).FieldByName("").IsValid())
		fmt.Println("not exist func:", reflect.ValueOf(s).MethodByName("").IsValid())

		m := map[int]int{}

		fmt.Println("not exist key:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
	}
	{
		fmt.Println()
		var a int
		typeOfA := reflect.TypeOf(a)
		aIns := reflect.New(typeOfA)
		fmt.Println(aIns.Type(), aIns.Kind())
	}
}

func add(a, b int) int {
	return a + b
}
func TestReflectFunc01(t *testing.T) {
	{
		funcVal := reflect.ValueOf(add)
		paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(24)}
		retList := funcVal.Call(paramList)
		fmt.Println("result:", retList[0].Int())
	}
	{
		fmt.Println()
		addSnp := func(a, b int) int {
			return a + b
		}
		funcVal := reflect.ValueOf(addSnp)
		paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(24)}
		retList := funcVal.Call(paramList)
		fmt.Println("result:", retList[0].Int())
	}
	{
		fmt.Println()
		funcVal := reflect.ValueOf(func(a, b int) int {
			return a + b
		})
		paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(24)}
		retList := funcVal.Call(paramList)
		fmt.Println("result:", retList[0].Int())
	}
}
