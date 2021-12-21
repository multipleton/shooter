package servers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/fail"
	"github.com/multipleton/shooter/server/http/utils"
)

type ServersController struct {
	serversService *ServersService
}

func (sc *ServersController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/servers", sc.All).Methods("GET")
	router.HandleFunc("/servers/users/{userId}/create", sc.Create).Methods("POST")
	router.HandleFunc("/servers/{serverId}/users/{userId}/join", sc.Join).Methods("POST")
	router.HandleFunc("/servers/{serverId}/users/{userId}/leave", sc.Leave).Methods("POST")
	router.HandleFunc("/servers/{serverId}/start", sc.Start).Methods("POST")
}

func (sc *ServersController) All(w http.ResponseWriter, r *http.Request) {
	result := sc.serversService.All()
	utils.Respond(w, http.StatusOK, result)
}

func (sc *ServersController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestParam("userId"))
		return
	}
	result, err := sc.serversService.CreateUserServer(userId)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, result)
}

func (sc *ServersController) Join(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serverId, err := strconv.Atoi(vars["serverId"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestParam("serverId"))
		return
	}
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestParam("userId"))
		return
	}
	err = sc.serversService.JoinUserServer(serverId, userId)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (sc *ServersController) Leave(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serverId, err := strconv.Atoi(vars["serverId"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestParam("serverId"))
		return
	}
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestParam("userId"))
		return
	}
	err = sc.serversService.LeaveUserServer(serverId, userId)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func (sc *ServersController) Start(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serverId, err := strconv.Atoi(vars["serverId"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestParam("serverId"))
		return
	}
	err = sc.serversService.StartServer(serverId)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func NewServersController(serversService *ServersService) *ServersController {
	return &ServersController{serversService: serversService}
}
