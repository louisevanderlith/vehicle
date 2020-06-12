package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Vehicles husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Vehicles: husk.NewTable(Vehicle{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Vehicles.Save()
}
