package auth

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) userGetHandler(ctx *gin.Context) {
	uid := ctx.Param("uid")

	auth, err := c.auth(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": true, "reason": err.Error()})
		return
	}

	user, err := auth.GetUser(ctx, uid)
	if err != nil {
		ctx.JSON(401, gin.H{"error": true, "reason": err.Error()})
		return
	}

	ctx.JSON(200, user)
}
