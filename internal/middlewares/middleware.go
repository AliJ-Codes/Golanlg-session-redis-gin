package middlewares

import (
	"net/http"
	"session-redis/internal/session"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		session_id, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		val, err := session.GetSession(rdb, session_id)
		if err != nil {
			if err != redis.Nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		err = session.UpdateTTL(rdb, session_id, time.Hour)
		c.Set("session_id", session_id)
		c.Set("user_id", val["user_id"])
		c.Set("role", val["role"])
		c.Next()
	}
}
