package auth

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func (c *Controller) auth(ctx *gin.Context) (*auth.Client, error) {
	return c.App.Auth(ctx)
}
