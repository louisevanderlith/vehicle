package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/vehicle/controllers"
)

func Setup(poxy resins.Epoxi) {
	//Vehicle
	vehCtrl := &controllers.VehicleController{}
	vehGroup := routing.NewRouteGroup("asset", mix.JSON)
	vehGroup.AddRoute("Create Vehicle", "", "POST", roletype.Owner, vehCtrl.Post)
	vehGroup.AddRoute("Vehicle by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, vehCtrl.GetByID)
	vehGroup.AddRoute("All Vehicles", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, vehCtrl.Get)
	poxy.AddGroup(vehGroup)
	/*ctrlmap := EnableFilter(s, host)

	vehCtrl := controllers.NewVehicleCtrl(ctrlmap)

	beego.Router("/v1/vehicle", vehCtrl, "post:Post")
	beego.Router("/v1/vehicle/:vehicleKey", vehCtrl, "get:GetByID")
	beego.Router("/v1/vehicle/all/:pagesize", vehCtrl, "get:Get")*/
}

/*
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
*/
