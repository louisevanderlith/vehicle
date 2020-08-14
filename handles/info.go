package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vehicle/core"
)

func SearchVehicle(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	result, err := core.GetVehicles(page, size)

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}

//:vehicleKey
func ViewVehicle(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

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

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	obj := core.Vehicle{}

	err := drx.JSONBody(r, &obj)

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

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}
