package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "tugasakhir/database"
	helpers "tugasakhir/helpers"
	models "tugasakhir/models"
)

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	photo.ID = uuid.New().String()
	userid, _ := c.Get("userid")
	photo.UserID = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Create(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to add item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Photo has been created successfully", "data": photo})
}

func ShowPhoto(c *gin.Context) {
	var photos []models.Photo
	userid, _ := c.Get("userid")

	database.DB.Where("userid = ?", userid).Find(&photos)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": photos})
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	photo.ID = c.Param("photoId")
	userid, _ := c.Get("userid")
	photo.UserID = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := database.DB.Model(&photo).Where("userid = ?", photo.UserID).Updates(&photo)

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
		return
	}

	database.DB.First(&photo, "id = ?", photo.ID)

	c.JSON(http.StatusOK, gin.H{"message": "data has been updated", "data": photo})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	photo.ID = c.Param("photoId")
	userid, _ := c.Get("userid")
	photo.UserID = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := database.DB.Where("userid = ?", photo.UserID).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Where("userid = ?", photo.UserID).Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"data": "Deleted successfully"})
}
