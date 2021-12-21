//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/core/engine"
	"github.com/multipleton/shooter/server/http"
	"github.com/multipleton/shooter/server/http/auth"
	"github.com/multipleton/shooter/server/http/models"
	"github.com/multipleton/shooter/server/http/servers"
	"github.com/multipleton/shooter/server/http/users"
	httpUtils "github.com/multipleton/shooter/server/http/utils"
	"github.com/multipleton/shooter/server/udp"
	"github.com/multipleton/shooter/server/utils"
)

type Application struct {
	HTTPServer *http.HTTPServer
	UDPServer  *udp.UDPServer
}

func composeHTTPControllers(
	authController *auth.AuthController,
	modelsController *models.ModelsController,
	serversController *servers.ServersController,
) *[]httpUtils.HTTPController {
	return &[]httpUtils.HTTPController{
		authController,
		modelsController,
		serversController,
	}
}

func Init(
	httpServerConfig http.HTTPServerConfig,
	udpServerConfig udp.UDPServerConfig,
) (*Application, error) {
	panic(wire.Build(
		mux.NewRouter,
		utils.NewEventEmitter,
		engine.NewManager,
		users.Module,
		servers.Module,
		auth.Module,
		models.Module,
		composeHTTPControllers,
		wire.Struct(new(http.HTTPServer), "config", "router", "controllers"),
		wire.Struct(new(udp.UDPServer), "config", "eventEmitter"),
		wire.Struct(new(Application), "HTTPServer", "UDPServer"),
	))
}
