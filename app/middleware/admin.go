package middleware

import (
	"ta_microservice_auth/app/models"
	"ta_microservice_auth/db"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("id")
		// if !ok {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		// 	return
		// }

		user := &models.Anggota{}
		err := db.Db.Preload("Role").First(user, userId).Error
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if user.Role == nil || user.Role.Name != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()

	}
}
