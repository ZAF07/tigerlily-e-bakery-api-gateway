package router

import (
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/api/rest/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
		// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins: true,
		// AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))


	/* Routes */
	inventory := r.Group("inventory")
	inventoryAPI := controller.NewInventoryAPI()

	{
		inventory.GET("", inventoryAPI.GetAllInventories)
	}

	return r
} 