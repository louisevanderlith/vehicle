package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/vehicle/core"
	"log"
	"net/http"
	"strconv"
)

func GetManufacturers(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(drx.FindParam(r, "year"))
	result, err := core.GetManufacturers(year)

	if err != nil {
		log.Println("Get Manufacturers Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var lst []string

	for name, _ := range result {
		lst = append(lst, name)
	}

	err = mix.Write(w, mix.JSON(lst))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(drx.FindParam(r, "year"))
	man := drx.FindParam(r, "manufacturer")

	result, err := core.GetModels(year, man)

	if err != nil {
		log.Println("Get Models Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var lst []string

	for name, _ := range result {
		lst = append(lst, name)
	}

	err = mix.Write(w, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}

func GetTrims(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(drx.FindParam(r, "year"))
	man := drx.FindParam(r, "manufacturer")
	mdl := drx.FindParam(r, "model")

	result, err := core.GetTrims(year, man, mdl)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var lst []string

	for name, _ := range result {
		lst = append(lst, name)
	}

	err = mix.Write(w, mix.JSON(lst))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
