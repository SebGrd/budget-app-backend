package controllers

import (
	"budget-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func GetExpenses(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	var expenses []models.Expense
	if err := db.Where("user_id = ?", userID).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func CreateExpense(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id").(uint)
	var payload struct {
		CategoryID uint    `json:"category_id" binding:"required" gorm:"not null"`
		Amount     float64 `json:"amount" binding:"required" gorm:"not null"`
		Date       string  `json:"date" binding:"required" gorm:"not null"`
		Note       string  `json:"note"`
		Recurring  bool    `json:"recurring" gorm:"default:false"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(payload.Date)
	expenseTime, err := time.Parse(`2006-01-02`, payload.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense := models.Expense{
		CategoryID: payload.CategoryID,
		Amount:     payload.Amount,
		Date:       expenseTime,
		Note:       payload.Note,
		Recurring:  payload.Recurring,
		UserID:     userID,
	}
	if err := db.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}
	c.JSON(http.StatusCreated, expense)
}

func UpdateExpense(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	var expense models.Expense
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context, db *gorm.DB) {
	userID := c.MustGet("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}
	if err := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Expense{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted"})
}
