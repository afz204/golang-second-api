package middleware

import (
	"api1/helper"
	"api1/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//validate the token user
func AuthorizeJWT(service service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		if headerAuth == "" {
			response := helper.BuildErrorResponse("Failed to process!", "Token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, err := service.ValidateToken(headerAuth)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim user:", claims["user_id"])
			log.Println("Claim issuer:", claims["issuer"])
		} else {
			response := helper.BuildErrorResponse("Token not valid!", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
