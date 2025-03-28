package routers

import (
	"github.com/gin-gonic/gin"
	controllers "jitD/controllers"
)

func BookRoutes(route *gin.Engine) {
	v1 := route.Group("v1/books")
	v1.GET("/", controllers.GetAllUser)
	v1.GET("/:id", controllers.GetBookById)
	v1.GET("/seller", controllers.GetSellerById)
	v1.POST("/", controllers.AddBook)
	v1.DELETE("/:id", controllers.DeleteBook)
}
