package middleware

import (
	"net/http"
	"strings"
	"ta_microservice_auth/app/controllers"
	"ta_microservice_auth/app/models"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
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
			errMsg := err.Error()
			res.Success = false
			res.Error = &errMsg
			c.JSON(401, res)
			c.Abort()
			return
		}

		_, found := claims["token_type"]
		if !found {
			errMsg := err.Error()
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if claims["token_type"] != "access_token" {
			errMsg := err.Error()
			res.Success = false
			res.Error = &errMsg
			c.JSON(401, res)
			c.Abort()
			return
		}
		// username, ok := claims["username"].(string)
		// if !ok {
		// 	errMsg := "Invalid username"
		// 	res.Success = false
		// 	res.Error = &errMsg
		// 	c.JSON(http.StatusUnauthorized, res)
		// 	c.Abort()
		// 	return
		// }

		// user, err := models.FindUserByUsername(db.Db, username)
		// if err != nil || user == nil {
		// 	errMsg := "User not found"
		// 	res.Success = false
		// 	res.Error = &errMsg
		// 	c.JSON(http.StatusUnauthorized, res)
		// 	c.Abort()
		// 	return
		// }

		// c.Set("user", user)
		c.Next()
	}

}
