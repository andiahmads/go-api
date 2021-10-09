package main

import (
	"github.com/andiahmads/go-api/config"
	controller "github.com/andiahmads/go-api/controllers"
	"github.com/andiahmads/go-api/middleware"
	"github.com/andiahmads/go-api/repository"
	service "github.com/andiahmads/go-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()

	//call service
	userService service.UserService = service.NewUserService(userRepository)
	authService service.AuthService = service.NewAuthService(userRepository)

	//call controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	r := gin.Default()
	authRouters := r.Group("api/auth")
	{
		authRouters.POST("/login", authController.Login)

		authRouters.POST("/register", authController.Register)
	}

	userRouters := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRouters.GET("/profile", userController.Profile)
		userRouters.PUT("/profile", userController.Update)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//testing
//sssdfdfdfdf

// //testing
// dfdf
