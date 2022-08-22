package ctrls

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pentag.kr/Meet-n-Auth-Server/models"
	"pentag.kr/Meet-n-Auth-Server/modules/jwtparse"
)

type logout struct {
	Token string `binding:"required"`
}

func LogOut(c *gin.Context) {

	var inputToken logout
	if err := c.ShouldBindJSON(&inputToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	tokenClaim, err := jwtparse.ParseToken(inputToken.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	var token models.Token
	models.DB.Model(&token).Where("token_id= ?", tokenClaim.TokenID).Update("expired", 1)
	c.JSON(http.StatusOK, "success")

}
