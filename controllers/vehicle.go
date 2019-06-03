package controllers

import (
	"encoding/json"
	"log"
	"net/http"

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

	req.Serve(http.StatusOK, nil, result)
}

//:vehicleKey
func (req *VehicleController) GetByID() {
	key, err := husk.ParseKey(req.Ctx.Input.Param(":vehicleKey"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetVehicle(key)

	if err != nil {
		log.Println(err)
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

func (req *VehicleController) Post() {
	obj := core.Vehicle{}

	json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	result, err := obj.Create()

	if err != nil {
		log.Println(err)
		req.Serve(http.StatusInternalServerError, err, nil)
	}

	req.Serve(http.StatusOK, nil, result)
}
