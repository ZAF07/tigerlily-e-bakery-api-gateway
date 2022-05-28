package inventory

import (
	"context"
	"net/http"

	wsClient "github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/manager/websocket-client"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
)

type SyncInventoryService struct {
	logger logger.Logger
}

func NewSyncInventoryService() *SyncInventoryService {
	return &SyncInventoryService{
		logger: *logger.NewLogger(),
	}
}

func WatchInventory(ctx context.Context) {

	// hub := wsClient.NewHub()
	// go hub.Run()
	// serveWs(hub, )
}

func serveWs(hub *wsClient.Hub, w http.ResponseWriter, r *http.Request) {
	return
}
