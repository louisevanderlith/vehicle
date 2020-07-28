package core

import (
	"errors"
	"github.com/louisevanderlith/husk"
)

type Vehicle struct {
	VINKey    husk.Key
	FullVIN   string
	Series    SeriesInfo
	Colour    string
	PaintNo   string
	Month     int
	Year      int
	Engine    Engine
	Gearbox   Gearbox
	BodyStyle string
	Doors     int
	Trim      string
	Extra     []string
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

	rec := ctx.Vehicles.Create(obj)

	if rec.Error != nil {
		return nil, rec.Error
	}

	ctx.Vehicles.Save()

	return rec.Record, nil
}
