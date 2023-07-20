package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("", indexHandler)
	router.GET("/specs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "specs.tmpl", gin.H{})
	})
	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.tmpl", gin.H{})
	})

	router.Run("localhost:8080")

}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
