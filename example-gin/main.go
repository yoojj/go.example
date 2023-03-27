package main 

import (
	"log"
	"net/http"
	//"github.com/gin-gonic/gin"

	"example-gin/config"
	"example-gin/router"
)



func init() {
	setting.Setup()
}

func main()  {
	log.Println("main()")

	routersInit := router.InitRouter()
	
	server := &http.Server{
		Handler: routersInit,
	}

	server.ListenAndServe()

}