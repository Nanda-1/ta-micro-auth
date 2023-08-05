package middleware

import (
	"net/http"
	"strings"
	"ta_microservice_auth/app/controllers"
	"ta_microservice_auth/app/models"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := models.JsonResponse{Success: true}

		if len(strings.Split(c.Request.Header.Get("Authorization"), " ")) != 2 {
			errorMsg := "invalid token"
			res.Success = false
			res.Error = &errorMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		accessToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]

		claims, err := controllers.DecodeToken(accessToken)
		if err != nil {
			// Token decoding error, return 401 Unauthorized
			errorMsg := err.Error()
			res.Success = false
			res.Error = &errorMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		role, found := claims["role_id"]
		if !found {
			errMsg := "role_id not found in token"
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		// Check the underlying type of role and convert to int if needed
		var roleInt int
		switch role.(type) {
		case float64:
			roleInt = int(role.(float64))
		case int:
			roleInt = role.(int)
		default:
			errMsg := "Invalid role_id type in token"
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		// Periksa apakah role_id = 1
		if roleInt != 1 {
			errMsg := "Hanyan Admin Yang Bisa Menambahkan Data"
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
