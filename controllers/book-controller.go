package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/entity"
	"github.com/andiahmads/go-api/helpers"
	service "github.com/andiahmads/go-api/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Pagination(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}

}

func (c *bookController) All(context *gin.Context) {
	//isi book dalam bentuk array
	var books []entity.Book = c.bookService.All()
	res := helpers.BuildResponse(true, "ok!", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(context *gin.Context) {
	//get data with paramaters ID
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book entity.Book = c.bookService.FindByID(id)
	if (book == entity.Book{}) {
		res := helpers.BuildErrorResponse("Data not Found", "No data with givend id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "ok!", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")

		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		fmt.Println("ini user =", convertedUserID)
		if err == nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	fmt.Println("ini user id = ", bookUpdateDTO)
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	//peta kan JWT TOKEN
	claims := token.Claims.(jwt.MapClaims)
	//ambil user id
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission to edit", "You not owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}

}

func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to get Id", "no params id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	//peta kan JWT TOKEN
	claims := token.Claims.(jwt.MapClaims)
	//ambil user id
	userID := fmt.Sprintf("%v", claims["user_id"])

	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helpers.BuildResponse(true, "Delete", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission to edit", "You not owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

// claims token user yang login
func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	//AMBIL TOKEN
	claims := aToken.Claims.(jwt.MapClaims)

	//print claims
	id := fmt.Sprintf("%v", claims["user_id"])

	return id
}

func (c *bookController) Pagination(context *gin.Context) {
	code := http.StatusOK
	pagination := helpers.GeneratePagination(context)
	response := c.bookService.AllbookWithPagination(context, pagination)
	if !response.Success {
		code = http.StatusBadRequest
	}
	context.JSON(code, response)
}
