package inventory

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	// "websocket-go/manager/hub"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return checkOrigin(r)
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	fmt.Println("READ PUMP")
	defer func() {
		fmt.Println("READ DEFER")
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {

		// Client's message/data ENTRY POINT (websocket method)
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				log.Printf("EWHAT 4444444444444444 ; %v", err)
			}
			// This loop only breaks when there is an error
			break
		}
		// INIT DB/CACHE CALL HERE TO FETCH/UPDATE LATEST INVENTORY STATUS
		log.Println("This is the incomming message ---> ", string(message))
		// Format the message/data and broadcast to hub to send to all active clients
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	fmt.Println("WRITE_PUMP")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		fmt.Println("WRITE DEFER")
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
// func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {

// 	// Should authenticate the request first
// 	/*
// 		Verify request is coming from a trusted host via r.Header
// 		Verify the JWT token given
// 	*/
// 	log.Println("This is the host: ", r.Header)
// 	// log.Println("This is the headers : ", )

// 	// Here i am upgrading the HTTP connection to a Websocket Protocol connection
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Creating a new client connection data structure
// 	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
// 	// Register a new connection to the hub
// 	client.Hub.Register <- client

// 	// Allow collection of memory referenced by the caller by doing all work in
// 	// new goroutines.
// 	go client.WritePump()
// 	go client.ReadPump()
// }

func checkOrigin(r *http.Request) bool {
	// Validate token
	token := r.URL.Query()
	if g, ok := token["token"]; ok {
		authToken := g[0]
		err := validateToken(authToken)
		if err != nil {
			return false
		}
	}
	return true
}

func validateToken(t string) error {
	if t != "1234" {
		log.Printf("Token %s does not match %s", t, "1234")
		return errors.New("invalid token")
	}
	return nil
}
