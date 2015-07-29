package main

import (
	"github.com/sarath.sp06/slacker/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/slackmessage", handlers.SlackMessagePostHandler)
	router.Run(":8080")
}
