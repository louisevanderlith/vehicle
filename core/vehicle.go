package core

import (
	"errors"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
	"github.com/louisevanderlith/vehicle/core/bodytype"
)

type Vehicle struct {
	VINKey    hsk.Key
	FullVIN   string
	Series    SeriesInfo
	Colour    string
	PaintNo   string
	Engine    Engine
	Gearbox   Gearbox
	BodyStyle bodytype.Enum
	Doors     int
	Extra     []string
	Spare     bool
	Service   bool
	Condition string
	Issues    string
	Mileage   int
}

func (m Vehicle) Valid() error {
	return validation.Struct(m)
}

func GetVehicles(page, pagesize int) (records.Page, error) {
	return ctx.Vehicles.Find(page, pagesize, op.Everything())
}

func GetVehicle(key hsk.Key) (hsk.Record, error) {
	return ctx.Vehicles.FindByKey(key)
}

func (obj Vehicle) Create() (hsk.Key, error) {
	if ctx.Vehicles.Exists(byFullVIN(obj.FullVIN)) {
		return nil, errors.New("vehicle VIN already exists")
	}

	return ctx.Vehicles.Create(obj)
}
