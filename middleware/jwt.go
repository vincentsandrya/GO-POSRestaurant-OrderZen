package middleware

import (
	"net/http"
	"os"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/display"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	UserId string = "user_id"
	Email  string = "email"
	RoleId string = "role_id"
)

func AuthorizeHandlerCookies() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response dto.ResponseMsg

		// Ambil token dari cookie
		tokenString, err := c.Cookie("authToken")
		if err != nil {
			response.Messageresp = display.ErrorBearerTokenInvalid.ErrorDisp()
			c.AbortWithStatusJSON(display.ErrorBearerTokenInvalid.CodeErr, response)
			return
		}

		// Deklarasikan klaim JWT
		claims := &dto.JWTClaims{}

		// Parse dan verifikasi token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			// Pastikan metode tanda tangan yang digunakan adalah HMAC
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, display.ErrorWrongCredentialsLogin
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})

		if err != nil {
			response.Messageresp = err.Error()
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Verifikasi klaim token
		if !token.Valid {
			response.Messageresp = display.ErrorUnathorized.ErrorDisp()
			c.AbortWithStatusJSON(display.ErrorUnathorized.CodeErr, response)
			return
		}

		c.Set(UserId, claims.UserId)
		c.Set(Email, claims.Email)
		c.Set(RoleId, claims.RoleId)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
