package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/vehicle/controllers"
)

func Setup(e resins.Epoxi) {
	vehCtrl := &controllers.Vehicles{}
	e.JoinBundle("/", roletype.Owner, mix.JSON, vehCtrl)
}
