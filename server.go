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
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	userRepository     repository.UserRepository     = repository.NewUserRepository(db)
	bookRepository     repository.BookRepository     = repository.NewBookRepository(db)
	categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(db)

	//call service
	jwtService      service.JWTService      = service.NewJWTService()
	userService     service.UserService     = service.NewUserService(userRepository)
	authService     service.AuthService     = service.NewAuthService(userRepository)
	bookService     service.BookService     = service.NewBookService(bookRepository)
	categoryService service.CategoryService = service.NewCategoryService(categoryRepository)

	//call controller
	authController     controller.AuthController     = controller.NewAuthController(authService, jwtService)
	userController     controller.UserController     = controller.NewUserController(userService, jwtService)
	bookController     controller.BookController     = controller.NewBookController(bookService, jwtService)
	categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
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

	bookRouters := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRouters.GET("/", bookController.All)
		bookRouters.GET("/withcategory", bookController.GetAllBookWithCategory)
		bookRouters.POST("/", bookController.Insert)
		bookRouters.GET("/:id", bookController.FindByID)
		bookRouters.PUT("/:id", bookController.Update)
		bookRouters.DELETE("/:id", bookController.Delete)
		bookRouters.GET("/pagination", bookController.Pagination)
	}
	categoryRouters := r.Group("api/category")
	{
		categoryRouters.POST("/", categoryController.Insert)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//testing
//sssdfdfdfdf

// //testing
// dfdf
