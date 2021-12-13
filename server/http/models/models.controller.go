package models

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/http/utils"
)

type ModelsController struct {
	modelsService *ModelsService
}

func (mc *ModelsController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/models/weapons", mc.Weapons).Methods("GET")
	router.HandleFunc("/models/items", mc.Items).Methods("GET")
	router.HandleFunc("/models/sides", mc.Sides).Methods("GET")
}

func (mc *ModelsController) Weapons(w http.ResponseWriter, r *http.Request) {
	result := mc.modelsService.GetAllWeapons()
	utils.Respond(w, http.StatusOK, result)
}

func (mc *ModelsController) Items(w http.ResponseWriter, r *http.Request) {
	result := mc.modelsService.GetAllItems()
	utils.Respond(w, http.StatusOK, result)
}

func (mc *ModelsController) Sides(w http.ResponseWriter, r *http.Request) {
	result := mc.modelsService.GetAllSides()
	utils.Respond(w, http.StatusOK, result)
}

func NewModelsController(modelsService *ModelsService) *ModelsController {
	return &ModelsController{modelsService: modelsService}
}
