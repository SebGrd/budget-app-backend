package main

import (
	"budget-go/controllers"
	"budget-go/middlewares"
	"budget-go/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	initDB()

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	setupRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Expense{},
		&models.Income{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func setupRoutes(router *gin.Engine) {
	router.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware)
	{
		authorized.GET("/profile", func(c *gin.Context) {
			controllers.Profile(c, db)
		})
		// CATEGORIES
		authorized.POST("/categories", func(c *gin.Context) {
			controllers.CreateCategory(c, db)
		})
		authorized.GET("/categories", func(c *gin.Context) {
			controllers.GetCategories(c, db)
		})
		authorized.PUT("/categories/:id", func(c *gin.Context) {
			controllers.UpdateCategory(c, db)
		})
		authorized.DELETE("/categories/:id", func(c *gin.Context) {
			controllers.DeleteCategory(c, db)
		})
		// EXPENSES
		authorized.POST("/expenses", func(c *gin.Context) {
			controllers.CreateExpense(c, db)
		})
		authorized.GET("/expenses", func(c *gin.Context) {
			controllers.GetExpenses(c, db)
		})
		authorized.PUT("/expenses/:id", func(c *gin.Context) {
			controllers.UpdateExpense(c, db)
		})
		authorized.DELETE("/expenses/:id", func(c *gin.Context) {
			controllers.DeleteExpense(c, db)
		})
	}
}
