package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vehicle/core"
)

type Vehicles struct {
}

func (req *Vehicles) Get(ctx context.Requester) (int, interface{}) {
	result := core.GetVehicles(1, 10)

	return http.StatusOK, result
}

//all/:pagesize
func (req *Vehicles) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()

	result := core.GetVehicles(page, size)

	return http.StatusOK, result
}

//:vehicleKey
func (req *Vehicles) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

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

func (req *Vehicles) Create(ctx context.Requester) (int, interface{}) {
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
