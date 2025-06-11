package main

import (
	"log"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/handlers"
	"loyaltea-server/internal/models"
	"loyaltea-server/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	// get the .env file
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

	userModel := models.NewUserModel(db.Database)
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

	offerModel := models.NewOfferModel(db.Database)
	offerService := services.NewOfferService(offerModel)
	offerHandler := handlers.NewOfferHandler(offerService)

	// offer routes
	router.POST("/offer/mailgun", offerHandler.ReceiveOffer)

	// //send a test email
	// apikey := os.Getenv("MAILGUN_API_KEY")
	// id, err := SendSimpleMessage("sandboxc4c89b92a012423a819e1762a284ab33.mailgun.org", apikey)
	// fmt.Println(id)
	// fmt.Println(err)

	log.Fatal(router.Run(":8080"))
}

// func SendSimpleMessage(domain, apiKey string) (string, error) {
// 	mg := mailgun.NewMailgun(domain, apiKey)
// 	//When you have an EU-domain, you must specify the endpoint:
// 	// mg.SetAPIBase("https://api.eu.mailgun.net")
// 	m := mg.NewMessage(
// 		"Mailgun Sandbox <postmaster@sandboxc4c89b92a012423a819e1762a284ab33.mailgun.org>",
// 		"Hello Keyur Bilgi",
// 		"Congratulations Keyur Bilgi, you just sent an email with Mailgun! You are truly awesome!",
// 		"Keyur Bilgi <platinumpizza29@gmail.com>",
// 	)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
// 	defer cancel()

// 	_, id, err := mg.Send(ctx, m)
// 	return id, err
// }
