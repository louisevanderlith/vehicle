package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"os"
	"reflect"
)

type context struct {
	Vehicles husk.Table
}

var ctx context

func CreateContext() {
	defer seed()
	ctx = context{
		Vehicles: husk.NewTable(Vehicle{}),
	}
}

func seed() {
	profiles, err := vehicleSeed()

	if err != nil {
		panic(err)
	}

	err = ctx.Vehicles.Seed(profiles)

	if err != nil {
		panic(err)
	}
}

func vehicleSeed() (collections.Enumerable, error) {
	f, err := os.Open("db/vehicles.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Vehicle
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func Shutdown() {
	ctx.Vehicles.Save()
}
