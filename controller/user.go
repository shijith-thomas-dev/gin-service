package controller

import (
	"gin-service/config"
	"gin-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	user := models.User{
		Name:     "shijith",
		Email:    "abc@abc.com",
		Password: "PASS",
	}
	// var user models.User
	// config.DB.Create(c.BindJSON(&user))
	config.DB.Create(&user)
	c.JSON(http.StatusAccepted, &user)

}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)

}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}
