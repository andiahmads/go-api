package middleware

import (
	"log"
	"net/http"

	"github.com/andiahmads/go-api/helpers"
	service "github.com/andiahmads/go-api/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {

		//get header token
		authHeader := c.GetHeader("Authorization")

		//cek auth header
		if authHeader == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)

		// jika token valid
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]:", claims["user_id"])
			log.Println("Claim[issuer]:", claims["issuer"])
		} else {
			log.Println(err)
			response := helpers.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}

}
