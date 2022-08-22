package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"pentag.kr/Meet-in-Auth-Server/ctrls"
	"pentag.kr/Meet-in-Auth-Server/models"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	r := gin.Default()
	models.ConnectDatabase()
	r.POST("/login", ctrls.Login)
	r.POST("/authenticate", ctrls.Authenticate)
	r.POST("/logout", ctrls.LogOut)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
