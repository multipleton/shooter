package servers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ServersController struct {
	serversService *ServersService
}

func (sc *ServersController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/servers/create", sc.Create).Methods("POST")
	router.HandleFunc("/servers/join", sc.Join).Methods("POST")
	router.HandleFunc("/servers/leave", sc.Join).Methods("POST")
}

func (sc *ServersController) Create(w http.ResponseWriter, r *http.Request) {

}

func (sc *ServersController) Join(w http.ResponseWriter, r *http.Request) {

}

func (sc *ServersController) Leave(w http.ResponseWriter, r *http.Request) {

}

func NewServersController(serversService *ServersService) *ServersController {
	return &ServersController{serversService: serversService}
}
