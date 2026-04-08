package utils
import (
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
)
type clientInfo struct {
	count    int
	lastSeen time.Time
}
var (
	mu      sync.Mutex
	clients = make(map[string]*clientInfo)
	limit   = 5
	window  = time.Minute
)
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string
		if userID, exists := c.Get("userID"); exists {
			key = "user:" + userID.(string)
		} else {
			key = "ip:" + c.ClientIP()
		}
		mu.Lock()
		cl, exists := clients[key]
		if !exists || time.Since(cl.lastSeen) > window {
			clients[key] = &clientInfo{count: 1, lastSeen: time.Now()}
			mu.Unlock()
			c.Next()
			return
		}
		cl.count++
		if cl.count > limit {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}
		mu.Unlock()
		c.Next()
	}
}