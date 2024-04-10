package grpc_client

import (
	"log"

	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Strategy GRPCClientInterface
	Conn     *grpc.ClientConn
}

// Returns a new instance of GRPCClient{}
func NewGRPCClient(port string) *GRPCClient {
	var conn *grpc.ClientConn
	conn, connErr := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
func (g *GRPCClient) SetStrategy(s GRPCClientInterface) {
	g.Strategy = s
}
