package udp

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/multipleton/shooter/server/udp/messages"
	"github.com/multipleton/shooter/server/utils"
)

type UDPServer struct {
	Config        UDPServerConfig
	connection    *net.UDPConn
	clients       map[string]*UDPClient
	EventEmitter  *utils.EventEmitter
	resultHandler *UDPResultHandler
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
	us.clients = make(map[string]*UDPClient)
	us.resultHandler = NewUDPResultHandler(us.send)
	us.EventEmitter.On(string(messages.STATE), us.resultHandler)
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
		buffer := make([]byte, 1024)
		n, addr, err := us.connection.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			continue
		}
		var message messages.UDPMessage
		err = json.Unmarshal(buffer[:n], &message)
		if err != nil {
			log.Println(err)
			continue
		}
		us.handle(addr, message)
	}
}

func (us *UDPServer) send(object interface{}) {
	data, err := json.Marshal(object)
	if err != nil {
		log.Printf("cannot convert object to JSON: %s", object)
	}
	for addrString := range us.clients {
		splittedAddr := strings.Split(addrString, ":")
		ip := net.ParseIP(splittedAddr[0])
		port, err := strconv.Atoi(splittedAddr[1])
		if err != nil {
			log.Printf("cannot resolve port of %s", addrString)
			continue
		}
		addr := net.UDPAddr{
			IP:   ip,
			Port: port,
		}
		_, err = us.connection.WriteToUDP(data, &addr)
		if err != nil {
			log.Printf("cannot send message to %s", addrString)
			log.Println(err)
			continue
		}
	}
}

func (us *UDPServer) handle(addr *net.UDPAddr, message messages.UDPMessage) {
	exists := false
	for _, messageType := range messages.Array {
		if messageType == message.MessageType {
			exists = true
			break
		}
	}
	if !exists {
		log.Printf("unrecognized message type: %s", message.MessageType)
		return
	}
	switch message.MessageType {
	case messages.CONNECT:
		us.addClient(addr, message.Data)
	case messages.DISCONNECT:
		us.deleteClient(addr)
	default:
		data := []interface{}{
			us.clients[addr.String()].Id,
			message.Data,
		}
		us.EventEmitter.Emit(string(message.MessageType), data)
	}
}

func (us *UDPServer) addClient(addr *net.UDPAddr, jsonObject utils.JSONObject) {
	var client UDPClient
	err := json.Unmarshal([]byte(jsonObject.Value), &client)
	if err != nil {
		log.Printf("cannot parse UDPClient object from %s", addr)
		log.Println(err)
		return
	}
	if us.clients[addr.String()] != nil {
		log.Printf("connection attempt from already connected client: %s", addr)
		return
	}
	us.clients[addr.String()] = &client
	log.Printf("client connected: %s", addr)
}

func (us *UDPServer) deleteClient(addr *net.UDPAddr) {
	delete(us.clients, addr.String())
	log.Printf("client disconnected: %s", addr)
}
