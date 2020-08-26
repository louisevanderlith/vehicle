package bodytype

type Enum = int

const (
	Bus Enum = iota
	Cabriolet
	Coupe
	DoubleCabBakkie
	Dropside
	Hatchback
	MPV
	PanelVan
	SUV
	Sedan
	SingleCabBakkie
	StationWagon
)

var vals = [...]string{
	"Bus",
	"Cabriolet",
	"Coupe",
	"Double Cab Bakkie",
	"Dropside",
	"Hatchback",
	"MPV",
	"Panel Van",
	"SUV",
	"Sedan",
	"Single Cab Bakkie",
	"Station Wagon",
}

func (e Enum) String() string {
	return vals[e]
}
