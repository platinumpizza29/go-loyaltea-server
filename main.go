// @title LoyalTea API
// @version 1.0
// @description Backend API for the LoyalTea mobile app
// @host localhost:8080
// @BasePath /
// @schemes http

package main

import (
	"log"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/handlers"
	"loyaltea-server/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	docs "loyaltea-server/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	offerModel := db.NewOfferModel(db.Database)
	offerService := services.NewOfferService(offerModel)
	offerHandler := handlers.NewOfferHandler(offerService)

	favModel := db.NewFavoriteModel(db.Database)
	favService := services.NewFavoriteService(favModel)
	favHandler := handlers.NewFavoriteHandler(favService)

	// shop and shopping plan models/services/handlers
	shopModel := db.NewShopModel(db.Database)
	shopService := services.NewShopService(shopModel)
	shopHandler := handlers.NewShopHandler(shopService)

	shoppingPlanModel := db.NewShoppingPlanModel(db.Database)
	shoppingPlanService := services.NewShoppingPlanService(shoppingPlanModel, userModel, shopModel)
	shoppingPlanHandler := handlers.NewShoppingPlanHandler(shoppingPlanService)

	// user routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// offer routes
	offerRoutes := router.Group("/offer")
	{
		offerRoutes.GET("/", offerHandler.GetOffers)
		offerRoutes.GET("/:id", offerHandler.GetOfferByID)
	}

	// fav routes
	favRoutes := router.Group("/fav")
	{
		favRoutes.POST("/", favHandler.CreateFav)
		favRoutes.GET("/:id", favHandler.GetFav)
		favRoutes.PUT("/:id", favHandler.UpdateFav)
		favRoutes.DELETE("/:id", favHandler.DeleteFav)
	}

	// shop routes
	shopRoutes := router.Group("/shops")
	{
		shopRoutes.GET("/", shopHandler.GetShops)
		shopRoutes.GET("/nearby", shopHandler.GetNearbyShops)
		shopRoutes.GET("/categories", shopHandler.GetCategories)
		shopRoutes.GET("/brand/:brand", shopHandler.GetShopsByBrand)
		shopRoutes.GET("/:id", shopHandler.GetShopByID)
		shopRoutes.POST("/", shopHandler.CreateShop)      // Admin only
		shopRoutes.PUT("/:id", shopHandler.UpdateShop)    // Admin only
		shopRoutes.DELETE("/:id", shopHandler.DeleteShop) // Admin only
		shopRoutes.POST("/batch", shopHandler.GetShopsBatchHandler)
	}

	// shopping plan routes
	planRoutes := router.Group("/planner")
	{
		planRoutes.POST("/", shoppingPlanHandler.CreatePlan)
		planRoutes.GET("/:id", shoppingPlanHandler.GetPlanByID)
		planRoutes.PUT("/:id", shoppingPlanHandler.UpdatePlan)
		planRoutes.DELETE("/:id", shoppingPlanHandler.DeletePlan)
		planRoutes.GET("/:id/progress", shoppingPlanHandler.GetPlanProgress)
		planRoutes.PUT("/:id/visit/:shopId", shoppingPlanHandler.MarkShopVisited)
		planRoutes.POST("/:id/shops", shoppingPlanHandler.AddShopToPlan)
		planRoutes.DELETE("/:id/shops/:shopId", shoppingPlanHandler.RemoveShopFromPlan)
		planRoutes.GET("/user/:id", shoppingPlanHandler.GetUserPlans)
		planRoutes.GET("/user/:id/active", shoppingPlanHandler.GetActivePlans)
		planRoutes.GET("/user/:id/completed", shoppingPlanHandler.GetCompletedPlans)
		planRoutes.GET("/user/:id/stats", shoppingPlanHandler.GetUserStats)
	}

	log.Fatal(router.Run(":8080"))
}
