package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()
	
	view := kong.ResourceMiddleware("vehicle.info.view", scrt, secureUrl, ViewVehicle)
	r.HandleFunc("/info/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware("vehicle.info.search", scrt, secureUrl, SearchVehicle)
	r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	//r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("vehicle.info.create", scrt, secureUrl, CreateVehicle)
	r.HandleFunc("/info", create).Methods(http.MethodPost)

	//update := kong.ResourceMiddleware("vehicle.info.update", scrt, secureUrl, )
	//r.HandleFunc("/info", update).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "vehicle.info.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst, //you service is available and allowed for this base url
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
