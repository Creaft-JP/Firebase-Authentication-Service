package auth

import (
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	App *firebase.App
}

func (c *Controller) Route(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/verify", c.verifyHandler)
	auth.GET("/users/:uid", c.userGetHandler)
}
