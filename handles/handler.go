package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewVehicle))).Methods(http.MethodGet)

	r.Handle("/info/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchVehicle))).Methods(http.MethodGet)
	//r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	r.Handle("/info", mw.Handler(http.HandlerFunc(CreateVehicle))).Methods(http.MethodPost)

	r.Handle("/lookup/manufacturers/{year:[0-9]+}", mw.Handler(http.HandlerFunc(GetManufacturers))).Methods(http.MethodGet)

	r.Handle("/lookup/models/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}", mw.Handler(http.HandlerFunc(GetModels))).Methods(http.MethodGet)

	r.Handle("/lookup/trim/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}/{model:[a-zA-Z]+}", mw.Handler(http.HandlerFunc(GetTrims))).Methods(http.MethodGet)

	/*
	//cars
		r.Handle("/cars", mw.Handler(http.HandlerFunc(GetCars))).Methods(http.MethodGet)
		r.Handle("/cars/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewCar))).Methods(http.MethodGet)
		r.Handle("/cars/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchCars))).Methods(http.MethodGet)
		r.Handle("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchCars))).Methods(http.MethodGet)
		r.Handle("/cars", mw.Handler(http.HandlerFunc(CreateCar))).Methods(http.MethodPost)
		r.Handle("/cars/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateCar))).Methods(http.MethodPut)
	*/

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "vehicle.info.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
