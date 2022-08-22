package ctrls

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pentag.kr/Meet-in-Auth-Server/models"
	"pentag.kr/Meet-in-Auth-Server/modules/cryption"
	"pentag.kr/Meet-in-Auth-Server/modules/jwtparse"
	"time"
)

const TokenExpireDuration = time.Hour * 720

type loginInput struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func Login(c *gin.Context) {
	var inputUser loginInput
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var user models.User
	result := models.DB.Take(&user, "Email = ? AND Password = ?", inputUser.Email, cryption.SHA512(inputUser.Password))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, "Username or PW is Wrong")
		return
	}
	token, err := createToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)

}

func createToken(userID uint64) (string, error) {
	// Create our own statement
	var lastTokenObject models.Token
	var nextTokenID uint64
	result := models.DB.Last(&lastTokenObject)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		nextTokenID = 1

	} else {
		nextTokenID = lastTokenObject.TokenID + 1

	}
	tokenObject := models.Token{TokenID: nextTokenID, Expired: false}
	result = models.DB.Create(&tokenObject)
	if result.Error != nil {
		panic(result.Error)
	}

	c := jwtparse.UserIDClaims{
		TokenID: nextTokenID,
		UserID:  userID, // Custom field
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // Expiration time
			Issuer:    "my-project",                               // Issuer
		},
	}
	// Creates a signed object using the specified signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Use the specified secret signature and obtain the complete encoded string token
	return token.SignedString(jwtparse.SecretKey)
}
