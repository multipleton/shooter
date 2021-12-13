package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AuthController struct {
	authService *AuthService
}

func (ac *AuthController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", ac.Login).Methods("POST")
	router.HandleFunc("/auth/logout", ac.Login).Methods("POST")
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {

}

func (ac *AuthController) Logout(w http.ResponseWriter, r *http.Request) {

}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}
