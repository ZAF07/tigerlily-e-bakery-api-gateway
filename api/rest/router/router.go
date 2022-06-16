package router

import (
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/api/rest/controller"

	i "github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/service/inventory"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
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

	inventoryAPI := controller.NewInventoryAPI(hub)
	checkoutAPI := controller.NewCheckoutAPI()

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
