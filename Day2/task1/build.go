package task1

import "github.com/gin-gonic/gin"

func Build() {
	router := gin.Default()

	router.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"Hello,gin Web Server",
		})
	})
	router.Run(":8080")

}
