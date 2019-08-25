package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vehicle/core"
)

type VehicleController struct {
}

//all/:pagesize
func (req *VehicleController) Get(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()

	result := core.GetVehicles(page, size)

	return http.StatusOK, result
}

//:vehicleKey
func (req *VehicleController) GetByID(ctx context.Contexer) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("vehicleKey"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetVehicle(key)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, rec
}

func (req *VehicleController) Post(ctx context.Contexer) (int, interface{}) {
	obj := core.Vehicle{}

	err := ctx.Body(&obj)

	if err != nil {
		return http.StatusBadRequest, err
	}

	result, err := obj.Create()

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}
