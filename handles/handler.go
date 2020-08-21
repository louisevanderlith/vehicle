package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	ins := kong.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	view := ins.Middleware("vehicle.info.view", scrt, ViewVehicle)
	r.HandleFunc("/info/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := ins.Middleware("vehicle.info.search", scrt, SearchVehicle)
	r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	//r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := ins.Middleware("vehicle.info.create", scrt, CreateVehicle)
	r.HandleFunc("/info", create).Methods(http.MethodPost)

	//update := ins.Middleware("vehicle.info.update", scrt, secureUrl, )
	//r.HandleFunc("/info", update).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "vehicle.info.view", scrt)

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
