package main 

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	//"example-gin/config/setting"
)



func init() {
	//setting.Setup()
}

func main()  {
    log.Println("main()")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"key": "val",
	  })
	})
	r.Run("localhost:8000")
}