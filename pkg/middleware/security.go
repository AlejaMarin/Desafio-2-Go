package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {

	requiredToken := os.Getenv("TOKEN")
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			//web.Failure(c, http.StatusUnauthorized, errors.New("Token No Encontrado"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token No Encontrado"})
			//c.Abort()
			return
		}
		if token != requiredToken {
			//web.Failure(c, http.StatusUnauthorized, errors.New("Token Inválido"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token Inválido"})
			//c.Abort()
			return
		}
		c.Next()
	}
}
