package utils

import (
	"budget-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsOwner(context *gin.Context, db *gorm.DB, userId interface{}) (models.User, error) {
	var err error
	contextUserId, _ := context.Get("user_id")
	var user models.User
	db.First(&user, userId)
	if user.ID == contextUserId {
		return user, err
	}
	err = fmt.Errorf("permission denied: user with ID %v is not the owner", contextUserId)
	return user, err
}
