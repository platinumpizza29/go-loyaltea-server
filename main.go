package main

import (
	"log"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/handlers"
	"loyaltea-server/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	//get the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBNAME := os.Getenv("DBNAME")
	DBURI := os.Getenv("DATABASE_URL")

	err = db.ConnectDB(DBURI, DBNAME)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// ping the server
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userModel := db.NewUserModel(db.Database)
	userService := services.NewUserService(userModel)
	userHandler := handlers.NewUserHandler(userService)

	// user routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	offerModel := db.NewOfferModel(db.Database)
	offerService := services.NewOfferService(offerModel)
	offerHandler := handlers.NewOfferHandler(offerService)

	offerRoutes := router.Group("/offer")
	{
		offerRoutes.GET("/", offerHandler.GetOffers)
		offerRoutes.GET("/:id", offerHandler.GetOfferByID)
	}

	router.GET("/offer/mailgun", offerHandler.VerifyWebhook)

	log.Fatal(router.Run(":8080"))
}
