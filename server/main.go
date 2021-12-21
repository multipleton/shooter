package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/multipleton/shooter/server/http"
	"github.com/multipleton/shooter/server/udp"
)

var httpServerConfig = http.HTTPServerConfig{
	Address: ":8080",
}

var udpServerConfig = udp.UDPServerConfig{
	Network: "udp4",
	Address: ":9000",
}

func main() {
	application, err := Init(httpServerConfig, udpServerConfig)
	if err != nil {
		log.Println("cannot initialize application")
		log.Fatalln(err)
	}
	go application.HTTPServer.Up()
	go application.UDPServer.Up()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	log.Println("interrupting...")
	application.HTTPServer.Down()
	application.UDPServer.Down()
}
