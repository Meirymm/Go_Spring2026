package utils
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role.(string) != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied: admin only"})
			return
		}
		c.Next()
	}
}