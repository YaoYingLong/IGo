package function

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Color struct {
	R, G, B int
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

func printMsg(msg *struct {
	id   int
	data string
}) {
	fmt.Println("msg", msg)
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b   int
	c   float32
	in1 int
	int
	innerS
}

func nestedStruct() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 20
	outer.innerS.in1 = 5
	outer.in2 = 10
	fmt.Println("outer1 :", *outer)
	outer2 := outerS{6, 7.5, 20, 60, innerS{5, 10}}
	fmt.Println("outer2 :", outer2)
}

type Wheel struct {
	Size int
}

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Wheel
	Engine
}

type Car2 struct {
	Wheel
	Engine2 struct {
		Power int
		Type  string
	}
}

func carFunc() {
	{
		c := Car{
			Wheel{
				Size: 18,
			},
			Engine{
				Type:  "1.4T",
				Power: 143,
			},
		}
		fmt.Println("car info:", c)
	}
	{
		c := Car{
			Wheel: Wheel{
				Size: 18,
			},
			Engine: Engine{
				Type:  "1.4T",
				Power: 143,
			},
		}
		fmt.Println("car info:", c)
	}
	{
		c := Car2{
			Wheel: Wheel{
				Size: 18,
			},
			Engine2: struct {
				Power int
				Type  string
			}{
				Type:  "1.4T",
				Power: 143,
			},
		}
		fmt.Println("car info:", c)
	}
	{
		c := Car2{
			Wheel{
				Size: 18,
			},
			struct {
				Power int
				Type  string
			}{
				Type:  "1.4T",
				Power: 143,
			},
		}
		fmt.Println("car info:", c)
	}
}

type Node struct {
	Data int
	Next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

func StructBase() {
	var point Point
	fmt.Println("ins before:", point)
	point.X = 15
	point.Y = 16
	fmt.Println("point before, X, Y:", point, point.X, point.Y)

	tank := new(Player)
	(*tank).Name = "Canon"
	tank.HealthPoint = 300
	fmt.Println("tank, X, Y:", *tank, tank.Name, tank.HealthPoint, tank.MagicPoint)

	tank2 := &Player{}
	(*tank2).Name = "Lisha"
	tank2.HealthPoint = 500
	fmt.Println("tank2, X, Y:", *tank2, tank2.Name, tank2.HealthPoint, tank2.MagicPoint)

	version := 1
	cmd := &Command{
		Name:    "version",
		Var:     &version,
		Comment: "show version",
	}
	fmt.Println("cmd, Name, Var, Comment:", *cmd, cmd.Name, *cmd.Var, cmd.Comment)

	cmd2 := Command{
		"version",
		&version,
		"show version",
	}
	fmt.Println("cmd2, Name, Var, Comment:", cmd2, cmd2.Name, *cmd2.Var, cmd2.Comment)

	ins := struct {
		Name    string
		Var     *int
		Comment string
	}{
		"version",
		&version,
		"show version",
	}
	fmt.Println("ins, Name, Var, Comment:", ins, ins.Name, *ins.Var, ins.Comment)

	msg := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	printMsg(msg)

	nestedStruct()

	carFunc()
}
