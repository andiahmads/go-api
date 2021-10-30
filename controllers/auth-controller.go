package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/entity"
	"github.com/andiahmads/go-api/helpers"
	service "github.com/andiahmads/go-api/services"
	"github.com/gin-gonic/gin"
)

//ctx digunakan untuk melihat response

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
	//masukkan service yang dibutuhkan disini
	authService service.AuthService
	jwtService  service.JWTService
}

//create new instance of authController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		refreshToken := c.jwtService.RefreshToken(strconv.FormatUint(v.ID, 10))
		v.Token = generateToken
		v.RefreshToken = refreshToken
		response := helpers.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Please Check Again email or Password", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Please cek your email", "duplicate email", helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		//create user
		createdUser := c.authService.CreateUser(registerDTO)
		fmt.Println(registerDTO)

		//generate Token
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helpers.BuildResponse(true, "ok!", createdUser)
		ctx.AbortWithStatusJSON(http.StatusCreated, response)
	}

}

func (c *authController) Logout(ctx *gin.Context) {

}
