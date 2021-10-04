package main

import (
	"github.com/andiahmads/go-api/config"
	controller "github.com/andiahmads/go-api/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()
	authRouters := r.Group("api/auth")
	{
		authRouters.POST("/login", authController.Login)

		authRouters.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//testing
//sssdfdfdfdf

// //testing
// dfdf
