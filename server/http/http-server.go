package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/http/utils"
)

type HTTPServer struct {
	Config      HTTPServerConfig
	Router      *mux.Router
	Controllers *[]*utils.HTTPController
	server      *http.Server
}

func (hs *HTTPServer) Up() {
	for _, controller := range *hs.Controllers {
		(*controller).HandleRoutes(hs.Router)
	}
	log.Printf("starting http server on %s", hs.Config.Address)
	hs.server = &http.Server{
		Addr:    hs.Config.Address,
		Handler: hs.Router,
	}
	err := hs.server.ListenAndServe()
	// TODO: investigate, it can be down if just got invalid body or smth like that
	if err != nil && err != http.ErrServerClosed {
		log.Println("http server down")
		log.Fatalln(err)
	}
}

func (hs *HTTPServer) Down() {
	if hs.server == nil {
		log.Println("cannot stop http server, is was not started")
	}
	log.Println("gracefully shuting down the http server")
	err := hs.server.Shutdown(context.Background())
	if err != nil {
		log.Println("http server gracefull shutdown failed")
		log.Fatalln(err)
	}
}
