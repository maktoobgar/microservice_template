package app

import (
	"fmt"
	"log"
	"net/http"

	g "service/gateway/global"
	"service/gateway/routes"

	"service/pkg/router"
)

func API() {
	// Print Info
	info()

	mux := new(router.Router)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", g.CFG.Gateway.IP, g.CFG.Gateway.Port),
		Handler: mux,
	}
	// Server uses ServeHTTP(ResponseWriter, *Request) method

	g.Server = server

	// Router Settings
	routes.HTTP(mux)

	// Run App
	log.Panic(server.ListenAndServe().Error())
}
