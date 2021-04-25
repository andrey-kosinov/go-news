package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/andrey.kosinov/go-news/auth"
)

// Middleware проверяющий аутентификацию пользователя
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session_name, _ := c.Cookie("session")
		session, ok := auth.CheckSession(session_name)
		if (ok) {
			c.Set("session",session)
			c.Next()
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.AbortWithStatus(503)
			return
		}
	}
}