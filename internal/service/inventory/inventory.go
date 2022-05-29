package inventory

import (
	"context"
	"log"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"google.golang.org/grpc"
)

type InventoryService struct {
	logs logger.Logger
	hubb *Hub
}

func NewInventoryService(h *Hub) *InventoryService {
	return &InventoryService{
		logs: *logger.NewLogger(),
		hubb: h,
	}
}

func (srv InventoryService) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {

	// Asynchronus API implementation
	// Create a connection instance and Dial the GRPC server
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(":8000", grpc.WithInsecure())
	if connErr != nil {
		srv.logs.ErrorLogger.Printf(" [SERVICE] Cannot connect to GRPC server")
		log.Fatalf("cannot connect to GRPC server : %+v", connErr)
	}
	defer conn.Close()

	// Initialises a new GRPC client service stub
	inventoryService := rpc.NewInventoryServiceClient(conn)

	resp, err = inventoryService.GetAllInventories(ctx, req)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error getting response from RPC server : %+v", err)
	}

	return
}

// serveWs handles websocket requests from the peer.
func (srv InventoryService) ServeWs(w http.ResponseWriter, r *http.Request) {

	// Should authenticate the request first
	/*
		Verify request is coming from a trusted host via r.Header
		Verify the JWT token given
	*/
	log.Println("This is the host: ", r.Header)
	// log.Println("This is the headers : ", )

	// Here i am upgrading the HTTP connection to a Websocket Protocol connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Creating a new client connection data structure
	client := &Client{Hub: srv.hubb, Conn: conn, Send: make(chan []byte, 256)}
	// Register a new connection to the hub
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
