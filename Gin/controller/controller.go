package controller

import (
	"fmt"
	dbs "Gin/database"
	model "Gin/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


var Db = dbs.InitDb()
	
func CreateUser(c *gin.Context) {
	fmt.Println("12345")
	var user model.Users
	body := model.Users{}
     err := c.BindJSON(&body) 
	 if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
    }
    user.ID= body.ID
    user.Name = body.Name
    user.Email= body.Email

    result := Db.Create(&user); 
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }

    c.JSON(http.StatusCreated, &user)
}


func GetUsers(c *gin.Context) {
	fmt.Println("get..")
	var users []model.User
	 result := Db.Find(&users); 
	 if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	fmt.Println("id")
	var user model.User
	result := Db.Where("id = ?",id).First(&user)
	fmt.Println(result)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":result.Error})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
	body := model.User{}

     err := c.BindJSON(&body); 
	  if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err})
			return
    }

    var user model.User

    if result := Db.First(&user, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }
    user.ID = body.ID

    Db.Save(&user)

    c.JSON(http.StatusOK,user)
}

func DeleteUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
    var user model.User
    result:=Db.Where("id = ?",id).Delete(&user)
	fmt.Println(result)
	 if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":result.Error })
        return
    }
    c.JSON(200, gin.H{"id":id})
}