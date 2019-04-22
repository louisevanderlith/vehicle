package core

import "github.com/louisevanderlith/husk"

type context struct {
	Vehicles husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Vehicles: husk.NewTable(new(Vehicle)),
	}
}

func Shutdown() {
	ctx.Vehicles.Save()
}
