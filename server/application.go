//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/multipleton/shooter/server/http"
	"github.com/multipleton/shooter/server/http/utils"
	"github.com/multipleton/shooter/server/udp"
)

type Application struct {
	HTTPServer *http.HTTPServer
	UDPServer  *udp.UDPServer
}

func Init(httpServerConfig http.HTTPServerConfig, udpServerConfig udp.UDPServerConfig) (*Application, error) {
	panic(wire.Build(
		mux.NewRouter,
		wire.Value(&[]*utils.HTTPController{}),
		wire.Struct(new(http.HTTPServer), "config", "router", "controllers"),
		wire.Struct(new(udp.UDPServer), "config"),
		wire.Struct(new(Application), "HTTPServer", "UDPServer"),
	))
}
