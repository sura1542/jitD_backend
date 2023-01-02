package main

import (
	// "fmt"
	// "context"
	// "net/http"

	"github.com/gin-gonic/gin"
	routes "jitD/routers"
)

func main() {

	// initail route
	router := gin.Default()

	// use middleware
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())
 
	// // provide route
	routes.BookRoutes(router)


	// configue on port 3000
	router.Run("0.0.0.0:3000")
}
