package ctrls

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pentag.kr/Meet-in-Auth-Server/modules/jwtparse"
)

var ctx = context.Background()

type authenticate struct {
	Token string `binding:"required"`
}

func Authenticate(c *gin.Context) {

	var inputToken authenticate
	if err := c.ShouldBindJSON(&inputToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	tokenClaim, err := jwtparse.ParseToken(inputToken.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, tokenClaim.UserID)
	return
}
