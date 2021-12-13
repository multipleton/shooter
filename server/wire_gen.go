// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/http"
	"github.com/multipleton/shooter/server/http/auth"
	"github.com/multipleton/shooter/server/http/models"
	"github.com/multipleton/shooter/server/http/servers"
	"github.com/multipleton/shooter/server/http/utils"
	"github.com/multipleton/shooter/server/udp"
)

// Injectors from application.go:

func Init(httpServerConfig2 http.HTTPServerConfig, udpServerConfig2 udp.UDPServerConfig) (*Application, error) {
	router := mux.NewRouter()
	authService := auth.NewAuthService()
	authController := auth.NewAuthController(authService)
	modelsService := models.NewModelsService()
	modelsController := models.NewModelsController(modelsService)
	serversService := servers.NewServersService()
	serversController := servers.NewServersController(serversService)
	v := composeHTTPControllers(authController, modelsController, serversController)
	httpServer := &http.HTTPServer{
		Config:      httpServerConfig2,
		Router:      router,
		Controllers: v,
	}
	udpServer := &udp.UDPServer{
		Config: udpServerConfig2,
	}
	application := &Application{
		HTTPServer: httpServer,
		UDPServer:  udpServer,
	}
	return application, nil
}

// application.go:

type Application struct {
	HTTPServer *http.HTTPServer
	UDPServer  *udp.UDPServer
}

func composeHTTPControllers(
	authController *auth.AuthController,
	modelsController *models.ModelsController,
	serversController *servers.ServersController,
) *[]utils.HTTPController {
	return &[]utils.HTTPController{
		authController,
		modelsController,
		serversController,
	}
}
