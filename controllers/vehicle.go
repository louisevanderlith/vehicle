package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/vehicle/core"
)

type VehicleController struct {
	control.APIController
}

func NewVehicleCtrl(ctrlMap *control.ControllerMap) *VehicleController {
	result := &VehicleController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

//all/:pagesize
func (req *VehicleController) Get() {
	page, size := req.GetPageData()

	result := core.GetVehicles(page, size)

	req.Serve(result, nil)
}

//:vehicleKey
func (req *VehicleController) GetByID() {
	key, err := husk.ParseKey(req.Ctx.Input.Param(":vehicleKey"))

	if err != nil {
		req.Serve(nil, err)
		return
	}

	req.Serve(core.GetVehicle(key))
}

func (req *VehicleController) Post() {
	obj := core.Vehicle{}

	json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	result, err := obj.Create()

	req.Serve(result, err)
}
