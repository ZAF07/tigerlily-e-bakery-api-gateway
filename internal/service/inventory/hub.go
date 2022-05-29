package inventory

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			fmt.Println("A NEW CONNECTION ESTABLISHED")
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				fmt.Println("DELETING A CLIENT")
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			fmt.Println("BROADCASTING A MESSAGE")
			for client := range h.Clients {
				select {
				case client.Send <- message:
					fmt.Println("SENDING A MESSAGE")
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
