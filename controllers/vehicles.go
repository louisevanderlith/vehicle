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

func (req *Vehicles) Get(c *gin.Context) {
	result := core.GetVehicles(1, 10)

	return http.StatusOK, result
}

//all/:pagesize
func (req *Vehicles) Search(c *gin.Context) {
	page, size := ctx.GetPageData()

	result := core.GetVehicles(page, size)

	return http.StatusOK, result
}

//:vehicleKey
func (req *Vehicles) View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

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

func (req *Vehicles) Create(c *gin.Context) {
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
