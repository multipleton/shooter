package udp

import (
	"log"
	"net"
)

type UDPServer struct {
	Config     UDPServerConfig
	connection *net.UDPConn
}

func (us *UDPServer) Up() {
	address, err := net.ResolveUDPAddr(us.Config.Network, us.Config.Address)
	if err != nil {
		log.Fatalln("cannot initialize udp address")
	}
	log.Printf("starting udp server on %s", us.Config.Address)
	connection, err := net.ListenUDP("udp4", address)
	if err != nil {
		log.Fatalln("cannot initialize udp connection")
	}
	us.connection = connection
	us.listen()
}

func (us *UDPServer) Down() {
	if us.connection == nil {
		log.Println("cannot stop udp server, it was not started")
	}
	log.Println("gracefully shuting down the udp server")
	err := us.connection.Close()
	if err != nil {
		log.Println("udp server gracefull shutdown failed")
		log.Fatalln(err)
	}
}

func (us *UDPServer) listen() {
	for {
		buf := make([]byte, 1024)
		n, address, err := us.connection.ReadFrom(buf)
		if err != nil {
			continue
		}
		log.Printf("\ngot %d from %s", n, address) // TODO: remove and add handler
	}
}
