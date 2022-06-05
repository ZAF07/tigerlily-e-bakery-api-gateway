package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/api/rest/router"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
)

func main() {
	fmt.Println("-** API GATEWAY STARTED **-")
	log := logger.NewLogger()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.ErrorLogger.Fatalf("[MAIN] Error connecting tcp port 8080: %+v\n", err)
	}

	// Start a new multiplexer passing in the main server
	m := cmux.New(l)
	httpListener := m.Match(cmux.HTTP1Fast())

	go serveHTTP(httpListener)

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		log.ErrorLogger.Fatalf("MUX ERR : %+v\n", err)
	}
}

func serveHTTP(l net.Listener) {
	h := gin.Default()
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		log.Fatalf("error serving HTTP : %+v", err)
	}
	logs := logger.NewLogger()
	logs.InfoLogger.Println("Started HTTP Server...")
	fmt.Println("HTTP Server Started ...")
}

//  NOT USED
// func serveWebsocket(l net.Listener) {
// 	// Initialise a new Websocket Client && start a go routine to listen for any events specifeid (Look in the hub.Run() for more details)
// 	hub := wsClient.NewHub()
// 	go hub.Run()

// 	// ENTRY POINT; Handler for websocket connections
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		// Handler to upgrade the HTTP connection to to a Websocket connection, create a new client connection, register the client to list of other connected clients
// 		wsClient.ServeWs(hub, w, r)
// 	})

// 	s := &http.Server{
// 		Addr: l.Addr().String(),
// 	}
// 	log.Println("SERVING WEBSOCKET")
// 	if err := s.Serve(l); err != nil {
// 		log.Fatalf("Error setting up websocket server : %v", err)
// 	}
// 	logs := logger.NewLogger()
// 	logs.InfoLogger.Println("Started HTTP Server...")
// }

// HTTP Server initialisation (using gin)
