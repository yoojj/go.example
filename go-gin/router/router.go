package router

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)



type auth struct {
	Username string
	Password string
}



func InitRouter() *gin.Engine {
	log.Println("InitRouter()")

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"key": "val",
		})
	})

	r.POST("/auth", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")

		user := auth{Username: username, Password: password}

		log.Println( user )
	})

	return r
}