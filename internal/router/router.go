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
	menuRepository := repository.NewMenuRepository(db)
	cartRepository := repository.NewCartRepository(db)

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

	driverService := service.NewDriverService(
		db,
		driverRepository,
	)

	customerService := service.NewCustomerService(
		db,
		customerRepository,
	)

	menuService := service.NewMenuService(
		db,
		menuRepository,
	)

	cartService := service.NewCartService(
		db,
		cartRepository,
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

	driverHandler := handler.NewDriverHandler(
		driverService,
	)

	customerHandler := handler.NewCustomerHandler(
		customerService,
	)

	menuHandler := handler.NewMenuHandler(
		menuService, restaurantService,
	)

	cartHandler := handler.NewCartHandler(
		cartService,
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
	protected.POST("/restaurant", restaurantHandler.CreateRestaurant) // no need to use

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
	protected.PATCH("/restaurants/:id/", restaurantHandler.UpdateRestaurantStatus)

	// =====================================
	// Driver Routes
	// =====================================

	// -------------------------------------
	// Protected Routes - DRIVER
	// -------------------------------------

	driver := protected.Group("/drivers")
	driver.Use(middleware.RequireRole("DRIVER"))

	{
		// Get driver profile
		driver.GET("/profile", driverHandler.GetDriverProfile)

		// Get driver details
		driver.GET("/me", driverHandler.GetDriverDetails)

		// Update driver profile
		driver.PATCH("/profile", driverHandler.UpdateDriverProfile)

		// Update driver status
		driver.PATCH("/status", driverHandler.UpdateDriverStatus)
	}

	// -------------------------------------
	// Protected Routes - ADMIN
	// -------------------------------------

	adminDrivers := protected.Group("/drivers")
	adminDrivers.Use(middleware.RequireRole("ADMIN"))

	{
		// Get all drivers
		adminDrivers.GET("", driverHandler.GetAllDrivers)

		// Get drivers by status
		adminDrivers.GET("/status", driverHandler.GetDriverByStatus)
	}

	// =====================================
	// Customer Routes
	// =====================================

	// -------------------------------------
	// Protected Routes - Customer
	// -------------------------------------
	customer := protected.Group("/customers")
	customer.Use(middleware.RequireRole("CUSTOMER"))

	{	
		// get customer profile
		customer.GET("/me", customerHandler.GetCustomerProfile)

		// get customer address
		customer.GET("/me/address", customerHandler.GetCostumerAddress)

		// update customer address
		customer.PATCH("/me/address", customerHandler.UpdateCustomerAddress)
	}

	// =====================================
	// Menu Routes
	// =====================================

	// -------------------------------------
	// Protected Routes - MERCHANT
	// -------------------------------------

	merchant := protected.Group("/restaurant/me")
	merchant.Use(middleware.RequireRole("MERCHANT"))
	{
		merchant.POST("/menus", menuHandler.CreateMenu)
		merchant.GET("/menus", menuHandler.GetMyMenu)
		merchant.GET("/menus/:id", menuHandler.GetMenuDetail)
		merchant.PATCH("/menus/:id", menuHandler.UpdateMenuInfo)
		merchant.PATCH("/menus/price/:id", menuHandler.UpdateMenuPrice)
		merchant.PATCH("/menus/availability/:id", menuHandler.UpdateMenuAvailability)
		merchant.PATCH("/menus/delete/:id", menuHandler.SoftDeleteMenu)
	}

	// -------------------------------------
	// Protected Routes - CUSTOMER MENU
	// -------------------------------------
	customerMenu := protected.Group("/restaurant/customer")
	customerMenu.Use(middleware.RequireRole("CUSTOMER"))
	{
		customerMenu.GET("/menus/:restaurantID", menuHandler.GetRestaurantMenus)
		customerMenu.GET("/menus/detail/:id", menuHandler.GetMenuDetail)
	}

	// example when other router must pass middleware auth
		// protected.GET("/users/me", userHandler.Me)

	// =====================================
	// Cart Routes
	// =====================================

	// -------------------------------------
	// Protected Routes - ADMIN
	// -------------------------------------
	adminCart := protected.Group("/admin/cart")
	adminCart.Use(middleware.RequireRole("ADMIN"))

	{
		adminCart.GET("/:customer_id", cartHandler.GetCartByCustomer)
	}

	// -------------------------------------
	// Protected Routes - CUSTOMER
	// -------------------------------------
	customerCart := protected.Group("/cart")
	customerCart.Use(middleware.RequireRole("CUSTOMER"))

	{
		customerCart.POST("", cartHandler.CreateCart)
		customerCart.GET("", cartHandler.GetMyCart)
		customerCart.DELETE("", cartHandler.DeleteCart)
	}

	return r
}