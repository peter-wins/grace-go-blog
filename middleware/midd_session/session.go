package midd_session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const KEY = "zcrblog2020"

/**
 * SESSION
 */
func MiddlewareSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("blogsession", store)
}
