package handlers

import (
	"net/http"
	"session-redis/internal/session"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func PanelHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.JSON(http.StatusOK, gin.H{
			"user_id": user_id,
			"role":    role,
		})
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "no",
				"err":    "invalid input",
			})
			c.Abort()
			return
		}
		if input.Password != "123" && input.Username != "admin" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "no",
				"error":  "username or password incorrect",
			})
			return
		}
		// TEST //
		user_id := 123123
		role := "admin"
		// END TEST //
		session_id, err := session.CreateSessionID()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = session.SetSession(rdb, session_id, time.Hour*24, user_id, role)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.SetCookie(
			"session_id",
			session_id,
			60*60*24*30,
			"/",
			"",
			false,
			true,
		)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
func LogoutHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		session_id, exists := c.Get("session_id")
		if !exists{
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}
		c.SetCookie(
			"session_id",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)
		delete, err := session.DeleteSession(rdb, session_id.(string))
		if err != nil || delete != 1{
			c.AbortWithStatus(http.StatusInternalServerError)
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
