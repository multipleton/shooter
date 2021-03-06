// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/core/engine"
	"github.com/multipleton/shooter/server/http"
	"github.com/multipleton/shooter/server/http/auth"
	"github.com/multipleton/shooter/server/http/models"
	"github.com/multipleton/shooter/server/http/servers"
	"github.com/multipleton/shooter/server/http/users"
	utils2 "github.com/multipleton/shooter/server/http/utils"
	"github.com/multipleton/shooter/server/udp"
	"github.com/multipleton/shooter/server/utils"
)

// Injectors from application.go:

func Init(httpServerConfig2 http.HTTPServerConfig, udpServerConfig2 udp.UDPServerConfig) (*Application, error) {
	router := mux.NewRouter()
	usersRepository := users.NewUsersRepository()
	usersService := users.NewUsersService(usersRepository)
	authService := auth.NewAuthService(usersService)
	authController := auth.NewAuthController(authService)
	modelsService := models.NewModelsService()
	modelsController := models.NewModelsController(modelsService)
	serversRepository := servers.NewServersRepository()
	eventEmitter := utils.NewEventEmitter()
	manager := engine.NewManager(eventEmitter)
	serversService := servers.NewServersService(serversRepository, usersService, manager)
	serversController := servers.NewServersController(serversService)
	v := composeHTTPControllers(authController, modelsController, serversController)
	httpServer := &http.HTTPServer{
		Config:      httpServerConfig2,
		Router:      router,
		Controllers: v,
	}
	udpServer := &udp.UDPServer{
		Config:       udpServerConfig2,
		EventEmitter: eventEmitter,
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
) *[]utils2.HTTPController {
	return &[]utils2.HTTPController{
		authController,
		modelsController,
		serversController,
	}
}
