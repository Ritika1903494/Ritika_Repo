package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
 )

func main() {
	router := gin.Default()
	router.Static("/asset", "./asset")
	router.LoadHTMLGlob("template/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "HTML Rendering",
		})
	})

	router.POST("/index", func(c *gin.Context) {
		file,err:= c.FormFile("file")

		if err!=nil{
			panic(err)
		}
		fmt.Println(file.Filename)

		err1:= c.SaveUploadedFile(file,"asset/"+file.Filename)
		if err1!=nil{
            panic(err1)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":8080")
}