package grpc_client

import (
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/constants"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Strategy GRPCClientInterface
	Conn     *grpc.ClientConn
}

// Returns a new instance of GRPCClient{}
func NewGRPCClient() *GRPCClient {
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if connErr != nil {
		logs := logger.NewLogger()
		logs.ErrorLogger.Printf(" [SERVICE] Cannot connect to GRPC server")
		log.Fatalf("cannot connect to GRPC server: %+v", connErr)
	}

	return &GRPCClient{
		Conn: conn,
	}
}

// Sets GRPCClient.Strategy()...
func (g *GRPCClient) SetStrategy(_type string) {
	switch _type {
	case constants.INVENTORY_SERVICE:
		g.Strategy = NewGRPCInventoryClient(g.Conn)
	}
}
