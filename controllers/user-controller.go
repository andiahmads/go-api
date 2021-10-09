package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/helpers"
	service "github.com/andiahmads/go-api/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) *userController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDto dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process Request", errDTO.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	//ambil data dari token
	claims := token.Claims.(jwt.MapClaims)
	// id, err := strconv.ParseUint(fmt.Sprintln("%v", claims["user_id"]), 10, 64)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDto.ID = id
	u := c.userService.Update(userUpdateDto)
	res := helpers.BuildResponse(true, "ok!", u)
	context.JSON(http.StatusOK, res)

}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helpers.BuildResponse(true, "ok!", user)
	context.JSON(http.StatusOK, res)

}
