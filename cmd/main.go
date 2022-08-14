package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/Tiger-Coders/tigerlily-bff/api/rest/router"
	"github.com/Tiger-Coders/tigerlily-bff/command"
	"github.com/Tiger-Coders/tigerlily-bff/config"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
)

func main() {
	log := logger.NewLogger()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.ErrorLogger.Fatalf("[MAIN] Error connecting tcp port 8080: %+v\n", err)
	}

	// ❌ NOT USED. THIS CAN BE DELETED. USED THIS TO MANUALLY CREATE A CONFIG FILE FOR INVENTORIES VIA THE CLI
	// Inject inventories into data file
	if cliErr := command.InjectInventoriesCmd.Execute(); cliErr != nil {
		log.ErrorLogger.Fatalf("Error Executing CLI commands : %+v\n", cliErr)
	}

	// Read inventories from the data file, pass to HTTP goroutine
	appConfig := config.InitAppConfig()

	// Start a new multiplexer passing in the main server
	m := cmux.New(l)
	httpListener := m.Match(cmux.HTTP1Fast())

	go serveHTTP(httpListener, appConfig)

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		log.ErrorLogger.Fatalf("MUX ERR : %+v\n", err)
	}
	fmt.Println("-** API GATEWAY STARTED **-")
}

/*
	TODO:
		SET UP READ/WRITE TIMEOUT
*/
func serveHTTP(l net.Listener, appConfig *config.AppConfig) {
	h := gin.Default()
	router.Router(h, appConfig)
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

func populateCache() {

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
