//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/http"
	"github.com/multipleton/shooter/server/http/auth"
	"github.com/multipleton/shooter/server/http/models"
	"github.com/multipleton/shooter/server/http/servers"
	"github.com/multipleton/shooter/server/http/utils"
	"github.com/multipleton/shooter/server/udp"
)

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

func Init(httpServerConfig http.HTTPServerConfig, udpServerConfig udp.UDPServerConfig) (*Application, error) {
	panic(wire.Build(
		mux.NewRouter,
		auth.Module,
		models.Module,
		servers.Module,
		composeHTTPControllers,
		wire.Struct(new(http.HTTPServer), "config", "router", "controllers"),
		wire.Struct(new(udp.UDPServer), "config"),
		wire.Struct(new(Application), "HTTPServer", "UDPServer"),
	))
}
