package server

type Server struct {
	address           string
	connectionHandler *ConnectionHandler
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
		connectionHandler: NewConnectionHandler(),
	}
}

func (server *Server) Run() {
	server.connectionHandler.run(server.address)
}