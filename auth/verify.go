package auth

import "github.com/gin-gonic/gin"

func (c *Controller) verifyHandler(ctx *gin.Context) {
	json := veryfyRequestContent{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{"error": true, "reason": err.Error()})
		return
	}

	auth, err := c.auth(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": true, "reason": err.Error()})
		return
	}

	token, err := auth.VerifyIDToken(ctx, json.IdToken)
	if err != nil {
		ctx.JSON(401, gin.H{"error": true, "reason": err.Error()})
		return
	}

	ctx.JSON(200, token)
}

type veryfyRequestContent struct {
	IdToken string `json:"idToken"`
}
