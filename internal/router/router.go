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

	// =====================================
	// Handler
	// =====================================

	authHandler := handler.NewAuthHandler(
		authService,
	)

	userHandler := handler.NewUserHandler(
		userService,
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

	// change user name
	protected.PATCH("/users/me/name", userHandler.ChangeUserName)

	// change user phone
	protected.PATCH("/users/me/phone", userHandler.ChangeUserPhone)

	// example when other router must pass middleware auth
		// protected.GET("/users/me", userHandler.Me)

	return r
}