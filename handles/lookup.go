package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/vehicle/core"
	"log"
	"net/http"
	"strconv"
)

func GetYears(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	result, err := core.GetYears()

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))
}

func GetManufacturers(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	year, _ := strconv.Atoi(ctx.FindParam("year"))
	result, err := core.GetManufacturers(year)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	year, _ := strconv.Atoi(ctx.FindParam("year"))
	man := ctx.FindParam("manufacturer")

	result, err := core.GetModels(year, man)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))
}

func GetTrims(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	year, _ := strconv.Atoi(ctx.FindParam("year"))
	man := ctx.FindParam("manufacturer")
	mdl := ctx.FindParam("model")

	result, err := core.GetTrims(year, man, mdl)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))
}
