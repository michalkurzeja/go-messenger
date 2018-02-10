package client

import (
	"os"
	"bufio"
)

var (
	input = make(chan string)
)

type Client struct {
	address string
	inputReader *InputReader
	connectionHandler *ConnectionHandler
}

func NewClient(address string) *Client {
	return &Client{
		address: 		   address,
		inputReader:       NewInputReader(bufio.NewScanner(os.Stdin)),
		connectionHandler: NewConnectionHandler(),
	}
}

func (client *Client) Run() {
	go client.inputReader.run(input)
	client.connectionHandler.run(client.address, input)
}


