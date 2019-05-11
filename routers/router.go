package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
	"github.com/louisevanderlith/secure/logic"
	"github.com/louisevanderlith/vehicle/controllers"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	vehCtrl := controllers.NewVehicleCtrl(ctrlmap)

	beego.Router("/v1/vehicle", vehCtrl, "post:Post")
	beego.Router("/v1/vehicle/:vehicleKey", vehCtrl, "get:GetByID")
	beego.Router("/v1/vehicle/all/:pagesize", vehCtrl, "get:Get")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := logic.NewMasterMap(s)

	emptyMap := make(core.ActionMap)
	emptyMap["GET"] = roletype.User
	emptyMap["POST"] = roletype.Owner

	ctrlmap.Add("/v1/vehicle", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap.ControllerMap
}
