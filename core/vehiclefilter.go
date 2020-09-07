package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type vehicleFilter func(obj Vehicle) bool

func (f vehicleFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(Vehicle))
}

func byYear(year int) vehicleFilter {
	return func(obj Vehicle) bool {
		return obj.Series.Year == year
	}
}

func byFullVIN(fullvin string) vehicleFilter {
	return func(obj Vehicle) bool {
		return obj.FullVIN == fullvin
	}
}
