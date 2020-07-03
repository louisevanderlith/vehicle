package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vehicle/core"
)

func SearchVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()
	result, err := core.GetVehicles(page, size)

	err = ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}

//:vehicleKey
func ViewVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.GetVehicle(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	obj := core.Vehicle{}

	err := ctx.Body(&obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := obj.Create()

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}
