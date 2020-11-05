package core

import (
	"errors"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
	"github.com/louisevanderlith/vehicle/core/bodytype"
	"time"
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

	Info          string `hsk:"size(128)"`
	Year          int    `orm:"null"`
	HasNatis      bool   `hsk:"default(false)"`
	EstValue      int64
	LicenseExpiry time.Time
}

func (o Vehicle) Valid() error {
	err := validation.Struct(o)
	if err != nil {
		return err
	}

	if o.Year > 0 && o.Year > time.Now().Year() {
		errors.New("year can't be in the future")
	}

	if o.Mileage < 0 {
		return errors.New("mileage can't be negative")
	}

	if o.HasNatis && o.LicenseExpiry.Before(time.Now()) {
		return errors.New("license has already expired")
	}

	return nil
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
