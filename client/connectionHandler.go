package client

import (
	"os"
	"net/url"
	"github.com/gorilla/websocket"
	"os/signal"
	"time"
)

type ConnectionHandler struct {
	interrupt chan os.Signal
}

func NewConnectionHandler() *ConnectionHandler {
	handler := &ConnectionHandler{}

	handler.interrupt = make(chan os.Signal, 1)
	signal.Notify(handler.interrupt, os.Interrupt)

	return handler
}

func (handler *ConnectionHandler) run(address string, channel <-chan string)  {
	connection := handler.connect(address)
	defer connection.Close()

	handler.loop(connection, channel)
}

func (handler *ConnectionHandler) connect(address string) *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}

	connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	handleError(err)

	return connection
}

func (handler *ConnectionHandler) loop(connection *websocket.Conn, channel <-chan string) {
	for {
		select {
		case message := <- channel:
			err := connection.WriteMessage(websocket.TextMessage, []byte(message))
			handleError(err)
		case <-handler.interrupt:
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			handleError(err)

			<-time.After(time.Second)

			return
		}
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}


