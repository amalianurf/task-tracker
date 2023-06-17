package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
				return
			} else {
				ctx.Redirect(http.StatusSeeOther, "/client/login")
				return
			}
		}

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{
					Error: err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "token invalid",
			})
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
