package core

import (
	"encoding/csv"
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"log"
	"os"
	"reflect"
	"strconv"
)

type VehicleContext interface {
	GetManufacturers(year int) (map[string]struct{}, error)
	GetModels(year int, manufacturer string) (map[string]struct{}, error)
	GetTrims(year int, manufacturer, model string) (map[string]struct{}, error)
}

func GetManufacturers(year int) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	err := ctx.Vehicles.Map(&result, Manufacturers(year))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetModels(year int, manufacturer string) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	err := ctx.Vehicles.Map(&result, Models(year, manufacturer))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTrims(year int, manufacturer, model string) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	err := ctx.Vehicles.Map(&result, Trim(year, manufacturer, model))

	if err != nil {
		return nil, err
	}

	return result, nil
}

type context struct {
	Vehicles husk.Table
}

var ctx context

func CreateContext() {
	//defer seed()
	ctx = context{
		Vehicles: husk.NewTable(Vehicle{}),
	}
}

func seed() {
	//profiles, err := vehicleSeed()
	profiles, err := vehicleDump()

	if err != nil {
		panic(err)
	}

	err = ctx.Vehicles.Seed(profiles)

	if err != nil {
		panic(err)
	}
}

func vehicleDump() (collections.Enumerable, error) {
	f, err := os.Open("db/vehicles.dump.tab")

	if err != nil {
		return nil, err
	}

	readr := csv.NewReader(f)

	readr.FieldsPerRecord = 6
	readr.Comma = '\t'
	readr.TrimLeadingSpace = true

	records, err := readr.ReadAll()
	log.Println("Records:", records)
	var result []Vehicle
	for i := 1; i < len(records); i++ {
		fields := records[i]

		year, err := strconv.Atoi(fields[3])

		if err != nil {
			return nil, err
		}

		veh := Vehicle{
			VINKey:  nil,
			FullVIN: "ABC",
			Series: SeriesInfo{
				Model:         fields[2],
				Manufacturer:  fields[1],
				AssemblyPlant: "",
				Month:         0,
				Year:          year,
				Trim:          "",
			},
			Colour:  "",
			PaintNo: "",
			Engine: Engine{
				Code:         "",
				SerialNo:     "",
				Output:       0,
				Fuel:         fields[5],
				Displacement: fields[4],
			},
			Gearbox: Gearbox{
				Code:     "",
				SerialNo: "",
				Gears:    0,
				Type:     "",
			},
			BodyStyle: 0,
			Doors:     0,
			Extra:     nil,
			Spare:     false,
			Service:   false,
			Condition: "",
			Issues:    "",
			Mileage:   0,
		}
		result = append(result, veh)

	}

	return collections.ReadOnlyList(reflect.ValueOf(result)), nil
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
