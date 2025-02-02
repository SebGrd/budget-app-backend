package controllers

import (
	"budget-go/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	var categories []models.Category
	if err := db.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)
	var payload struct {
		Name string `json:"name" binding:"required"`
		Icon string `json:"icon" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := models.Category{
		Name:   payload.Name,
		Icon:   payload.Icon,
		UserID: userID,
	}
	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	if err := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
