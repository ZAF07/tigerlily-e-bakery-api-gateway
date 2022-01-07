package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/api/rest/router"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API GATEWAY")
	l, err := net.Listen("tcp", ":8080")
if err != nil {
	fmt.Println("Something went wrong in the server startup")
	log.Fatalf("Error connecting tcp port 8000")
}
// logs.InfoLogger.Println("Successfull server init")

	h := gin.Default()
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}

	s.Serve(l)
}