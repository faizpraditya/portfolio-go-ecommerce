package routes

import (
	"github.com/faizpraditya/go-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	// just goin to search a keyword, no filtering etc in this project (at least initially)
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
