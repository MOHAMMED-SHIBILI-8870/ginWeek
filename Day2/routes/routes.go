package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		name := c.Query("name")
		if name == "" {
			name = "name"
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "helooo  " + name,
		})
	})

	r.Run(":8080")
}
