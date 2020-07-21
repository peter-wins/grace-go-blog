package midd_auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("username") == nil {
			c.Abort()
			c.Redirect(http.StatusFound, "/login")
		}
	}
}
