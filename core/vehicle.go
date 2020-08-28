package core

import (
	"errors"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vehicle/core/bodytype"
)

type Vehicle struct {
	VINKey    husk.Key
	FullVIN   string
	Series    SeriesInfo
	Colour    string
	PaintNo   string
	Engine    Engine
	Gearbox   Gearbox
	BodyStyle bodytype.Enum
	Doors     int
	Extra     []string
	Spare	bool
	Service	bool
	Condition	string
	Issues	string
	Mileage	int
}

func (m Vehicle) Valid() error {
	return husk.ValidateStruct(m)
}

func GetVehicles(page, pagesize int) (husk.Collection, error) {
	return ctx.Vehicles.Find(page, pagesize, husk.Everything())
}

func GetVehicle(key husk.Key) (husk.Recorder, error) {
	return ctx.Vehicles.FindByKey(key)
}

func (obj Vehicle) Create() (husk.Recorder, error) {
	if ctx.Vehicles.Exists(byFullVIN(obj.FullVIN)) {
		return nil, errors.New("vehicle VIN already exists")
	}

	rec, err := ctx.Vehicles.Create(obj)

	if err != nil {
		return nil, err
	}

	err = ctx.Vehicles.Save()

	if err != nil {
		return nil, err
	}

	return rec, nil
}
