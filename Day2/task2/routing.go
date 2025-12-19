package task2

import "github.com/gin-gonic/gin"

func Routing() {

	router := gin.Default()

	router.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"home":"welcome to Home!",
		})
	})

	router.GET("/about",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"about":"welcome to about page",
		})
	})
	router.GET("/contact",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"contact":"welcome to about page",
		})
	})
	router.Run(":8080")

}