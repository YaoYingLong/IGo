package function

import "fmt"

type Point struct {
	X int
	Y int
}

type Color struct {
	R, G, B byte
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

func StructBase() {
	var ins Point
	fmt.Println("ins before:", ins)
	ins.X = 15
	ins.Y = 16
	fmt.Println("ins before, X, Y:", ins, ins.X, ins.Y)

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
}
