package jwtparse

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"os"
	"pentag.kr/Meet-in-Auth-Server/models"
)

var SecretKey = []byte(os.Getenv("ACCESS_SECRET"))

func ParseToken(tokenString string) (*UserIDClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &UserIDClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserIDClaims); ok && token.Valid { // Verification token
		var tokenData models.Token
		result := models.DB.Take(&tokenData, "token_id = ?", claims.TokenID)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid token")
		}
		if tokenData.Expired == true {
			return nil, errors.New("token expired by user")
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

type UserIDClaims struct {
	TokenID uint64 `json:tokenid`
	UserID  uint64 `json:"userid"`
	jwt.StandardClaims
}
