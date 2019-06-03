package routers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
	"github.com/louisevanderlith/vehicle/controllers"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilter(s, host)

	vehCtrl := controllers.NewVehicleCtrl(ctrlmap)

	beego.Router("/v1/vehicle", vehCtrl, "post:Post")
	beego.Router("/v1/vehicle/:vehicleKey", vehCtrl, "get:GetByID")
	beego.Router("/v1/vehicle/all/:pagesize", vehCtrl, "get:Get")
}

func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(core.ActionMap)
	emptyMap["GET"] = roletype.User
	emptyMap["POST"] = roletype.Owner

	ctrlmap.Add("/v1/vehicle", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
