package server

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

type ConnectionHandler struct {
	upgrader *websocket.Upgrader
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		upgrader: &websocket.Upgrader{},
	}
}

func (handler *ConnectionHandler) run(address string) {
	http.HandleFunc("/", handler.listen)
	http.ListenAndServe(address, nil)

}

func (handler *ConnectionHandler) listen(writer http.ResponseWriter, request *http.Request) {
	connection := handler.openConnection(writer, request)

	for {
		messageType, message, err := connection.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("[%d] %s", messageType, message)
	}
}

func (handler *ConnectionHandler) openConnection(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	connection, err := handler.upgrader.Upgrade(writer, request, nil)
	handleError(err)

	return connection
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}