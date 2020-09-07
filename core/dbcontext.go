package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Vehicles husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Vehicles: husk.NewTable(Vehicle{}),
	}
}

func Shutdown() {
	ctx.Vehicles.Save()
}
