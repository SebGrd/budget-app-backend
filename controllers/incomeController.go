package controllers

import (
	"budget-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetIncomes(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	var incomes []models.Income
	if err := db.Where("user_id = ?", userID).Find(&incomes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch incomes"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, incomes)
}

func CreateIncome(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)
	var payload struct {
		Amount    float64 `json:"amount" binding:"required" gorm:"not null"`
		Date      string  `json:"date" binding:"required" gorm:"not null"`
		Note      string  `json:"note"`
		Recurring bool    `json:"recurring" gorm:"default:false"`
		Source    string  `json:"source" gorm:"not null"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	income := models.Income{
		Amount:    payload.Amount,
		Date:      payload.Date,
		Note:      payload.Note,
		Recurring: payload.Recurring,
		Source:    payload.Source,
		UserID:    userID,
	}
	if err := db.Create(&income).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create income"})
		return
	}
	c.JSON(http.StatusCreated, income)
}

func UpdateIncome(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid income ID"})
		return
	}

	var income models.Income
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&income).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Income not found"})
		return
	}

	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&income).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update income"})
		return
	}

	c.JSON(http.StatusOK, income)
}

func DeleteIncome(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid income ID"})
		return
	}
	if err := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Income{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete income"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Income deleted"})
}
