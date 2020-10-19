package main

import (
	"flag"
	"github.com/louisevanderlith/vehicle/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/vehicle/core"
)

func main() {
	securty := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	manager := flag.String("manager", "http://localhost:8097", "Manager Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8098",
		Handler:      handles.SetupRoutes(*srcSecrt, *securty, *manager),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
