package server

type Client interface {
	ID() int
}

type RssServer struct {
	clients []Client
}

func NewRssServer() RssServer {
	return RssServer{clients: make([]Client, 0)}
}

// Register client to receive messages from RssServer.
// Client with the same ID will not be registered once again.
func (s *RssServer) Register(c Client) {
	if !s.IsRegistered(c) {
		s.clients = append(s.clients, c)
	}
}

func (s RssServer) IsRegistered(c Client) bool {
	for _, e := range s.clients {
		if e.ID() == c.ID() {
			return true
		}
	}
	return false
}
