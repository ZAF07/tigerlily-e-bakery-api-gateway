package router

import (
	"github.com/Tiger-Coders/tigerlily-bff/api/rest/controller"
	"github.com/Tiger-Coders/tigerlily-bff/config"
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	i "github.com/Tiger-Coders/tigerlily-bff/internal/service/inventory"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, appConfig *config.AppConfig) *gin.Engine {
	log := logger.NewLogger()
	log.InfoLogger.Println("[ROUTER] ROUTER HAS RECEIVED THE IDENTIFIERS OF INVENTORIES", appConfig)
	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins:  true,
		// ❌ Might want to only allow specific host for security
		// AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	/* Routes */
	inventory := r.Group("inventory")
	checkout := r.Group("checkout")

	// ❌ TODO: This could be moved to main.go
	hub := i.NewHub()
	go hub.Run()

	// THIS HAS TO CHANGE TO APPCONFIG NOT INVS
	inventoryAPI := controller.NewInventoryAPI(hub, appConfig)
	checkoutAPI := controller.NewCheckoutAPI(appConfig)

	{
		/* INVENTORY API */
		inventory.GET("", inventoryAPI.GetAllInventories)
		inventory.GET("/cache", inventoryAPI.GetAllInventoriesCache)
		inventory.GET("/ws", inventoryAPI.WsInventory)

		/* CHECKOUT API */
		checkout.POST("", checkoutAPI.Checkout)
	}

	return r
}
