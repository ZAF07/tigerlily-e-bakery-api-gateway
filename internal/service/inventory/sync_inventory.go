package inventory

import (
	"context"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/manager/hub"
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

}

func serveWs(hub *hub.Hub, w http.ResponseWriter, r *http.Request) {
	return
}
