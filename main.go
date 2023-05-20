package main

import (
	"awesomeProject1/retailshift_manager"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	fmt.Println("INIT")
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.Default()

	r.Any("/healthcheck", func(c *gin.Context) {
		c.Status(200)
	})

	r.POST("/retailshift/create", func(c *gin.Context) {
		retailshift_manager.OnHandleCreateRetailshiftController(c)
	})

	r.POST("/retailshift/update", func(c *gin.Context) {
		retailshift_manager.OnHandleUpdateRetailshiftController(c)
	})

	r.Run(fmt.Sprintf(":%s", port))
}
