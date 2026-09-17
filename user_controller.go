package controllers

import (
	"10env/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	result := uc.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(201, user)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	result := uc.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	c.JSON(200, users)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	result := uc.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(500, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var request UpdateUserRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	var user models.User

	result := uc.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(500, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	user.Name = request.Name

	result = uc.DB.Save(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	result := uc.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(500, gin.H{
			"error": "Failed to find user",
		})
		return
	}

	result = uc.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User deleted successfully",
	})
}
