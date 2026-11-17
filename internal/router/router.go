package router

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kholiklutfi29/nourish-dispatch/internal/config"
	"github.com/kholiklutfi29/nourish-dispatch/internal/handler"
	"github.com/kholiklutfi29/nourish-dispatch/internal/jwt"
	"github.com/kholiklutfi29/nourish-dispatch/internal/middleware"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
	"github.com/kholiklutfi29/nourish-dispatch/internal/service"
)

func SetupRouter(db *sql.DB) *gin.Engine {

	// =====================================
	// JWT
	// =====================================

	jwtService := jwt.NewJWTService(
		config.GetJWTSecret(),
		24 * time.Hour, 
	)

	// =====================================
	// Repository
	// =====================================

	userRepository := repository.NewUserRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	driverRepository := repository.NewDriverRepository(db)
	restaurantRepository := repository.NewRestaurantRepository(db)

	// =====================================
	// Service
	// =====================================

	profileService := service.NewProfileService(
		customerRepository,
		driverRepository,
		restaurantRepository,
	)

	authService := service.NewAuthService(
		db,
		userRepository,
		profileService,
		jwtService,
	)

	userService := service.NewUserService(
		db,
		userRepository,
	)

	restaurantService := service.NewRestaurantService(
		db,
		restaurantRepository,
	)

	// =====================================
	// Handler
	// =====================================

	authHandler := handler.NewAuthHandler(
		authService,
	)

	userHandler := handler.NewUserHandler(
		userService,
	)

	restaurantHandler := handler.NewRestaurantHandler(
		restaurantService,
	)

	// =====================================
	// Gin
	// =====================================

	r := gin.New()

	// =====================================
	// Global Middleware
	// =====================================

	// all request will pass this 3 middleware
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
	)

	// =====================================
	// API
	// =====================================

	api := r.Group("/api/v1")

	// =====================================
	// Public Routes
	// =====================================

	auth := api.Group("/auth")

	auth.POST(
		"/register",
		authHandler.Register,
	)

	auth.POST(
		"/login",
		authHandler.Login,
	)

	// =====================================
	// Protected Routes
	// =====================================

	// must pass auth middleware token verification
	protected := api.Group("")

	protected.Use(
		middleware.Auth(jwtService),
	)

	// =====================================
	// Users Routes
	// =====================================

	// change user name
	protected.PATCH("/users/me/name", userHandler.ChangeUserName)

	// change user phone
	protected.PATCH("/users/me/phone", userHandler.ChangeUserPhone)

	// change password
	protected.PATCH("/users/me/password", userHandler.ChangeUserPassword)

	// =====================================
	//  Restaurant Routes
	// =====================================

	// create new restaurant
	protected.POST("/restaurant", restaurantHandler.CreateRestaurant)

	// get all restaurant
	api.GET("/restaurants", restaurantHandler.GetAllRestaurants)

	// get restaurant by id
	api.GET("/restaurants/:id", restaurantHandler.GetRestaurantByID)

	// soft delete
	protected.DELETE("/restaurants/:id", restaurantHandler.SoftDeleteRestaurant)

	// update restaurant
	protected.PATCH("/restaurants/:id", restaurantHandler.UpdateRestaurant)

	// get owner restaurant
	protected.GET("/restaurants/me", restaurantHandler.GetMyRestaurants)

	// update restaurant status
	protected.PATCH("restaurants/:id/", restaurantHandler.UpdateRestaurantStatus)

	// example when other router must pass middleware auth
		// protected.GET("/users/me", userHandler.Me)

	return r
}