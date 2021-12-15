package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/fail"
	"github.com/multipleton/shooter/server/http/utils"
)

type AuthController struct {
	authService *AuthService
}

func (ac *AuthController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", ac.Login).Methods("POST")
	router.HandleFunc("/auth/logout", ac.Logout).Methods("POST")
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var dto LoginAuthDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestBody())
	}
	result, err := ac.authService.Login(dto)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, result)
}

func (ac *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	var dto LogoutAuthDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, fail.InvalidRequestBody())
	}
	err = ac.authService.Logout(dto)
	fmt.Println(err)
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, nil)
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}
