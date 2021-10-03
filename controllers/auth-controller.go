package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ctx digunakan untuk melihat response

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	//masukkan service yang dibutuhkan disini

}

//create new instance of authController
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello login",
	})

}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello register",
	})
}
