package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vehicle/core"
)

type VehicleController struct {
	xontrols.APICtrl
}

//all/:pagesize
func (req *VehicleController) Get() {
	page, size := req.GetPageData()

	result := core.GetVehicles(page, size)

	req.Serve(http.StatusOK, nil, result)
}

//:vehicleKey
func (req *VehicleController) GetByID() {
	key, err := husk.ParseKey(req.FindParam("vehicleKey"))

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

	err := req.Body(&obj)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result, err := obj.Create()

	if err != nil {
		log.Println(err)
		req.Serve(http.StatusInternalServerError, err, nil)
	}

	req.Serve(http.StatusOK, nil, result)
}
