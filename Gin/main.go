package main


import (
	"github.com/gin-gonic/gin"
	control "Gin/controller"
 )

 func main() {
	r := gin.Default()
	r.GET("/Main", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello gin",
		})
	})
	r.POST("/Main/CreateUser",control.CreateUser)
	r.GET("/Main/GetUsers",control.GetUsers)
	r.GET("/Main/GetUser/:id",control.GetUser)
	r.PATCH("/Main/UpdateUser/:id",control.UpdateUser)
	r.DELETE("/Main/DeleteUser/:id",control.DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}